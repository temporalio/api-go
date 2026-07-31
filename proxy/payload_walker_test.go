package proxy

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/update/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
)

const (
	payloadFullName  = "temporal.api.common.v1.Payload"
	payloadsFullName = "temporal.api.common.v1.Payloads"
	anyFullName      = "google.protobuf.Any"
	temporalPrefix   = "temporal.api."

	// maxWalkDepth bounds how far the walker will descend. If there is a cycle
	// (e.g. A holds a B, and a B holds an A) we hit this max depth and fail
	// rather than panic with a stack overflow.
	maxWalkDepth = 50
)

// PayloadWalker verifies that VisitPayloads can reach every payload that a proto
// message tree is capable of holding, no matter how deeply it is nested.
//
// It works by walking the descriptors of a root message and populating exactly
// one leaf at a time: for each field it creates a fresh child message, descends
// into it, then clears the field again. Whenever the walk reaches a payload
// container (Payload, Payloads, or an Any stuffed with a payload-bearing
// message) it runs VisitPayloads against the *root* message and asserts the
// Visitor fires exactly once, proving the generated visitor code has a path
// from the root all the way down to that leaf.
//
// The zero value is not usable; construct one with RunPayloadWalker.
type PayloadWalker struct {
	t *testing.T

	// MessagesVisited is the number of proto messages the walker descended into.
	// Non-Temporal messages (other than Any) are not walked and not counted.
	MessagesVisited int

	// PayloadsSeen has one entry per payload container the walker reached, in
	// discovery order, formatted as "<field path> [<message type>]".
	PayloadsSeen []string

	// Skipped has one entry per subtree the walker declined to enter, along with
	// the reason. (e.g. self-referential fields such as Failure.cause would
	//recurse forever.)
	Skipped []string

	root proto.Message
	// rootName is the full name of the message the walk started from. Kept
	// separately from path[0] so that Report stays correct no matter where in the
	// walk it is called from.
	rootName string

	// path is the field path from root down to the message currently being walked.
	path []string

	// pendingVisits counts payload containers that have been populated and
	// handed to VisitPayloads but that the Visitor has not reported back yet.
	// Because only one payload is populated at a time it is 1 for the duration of
	// a single VisitPayloads call and 0 everywhere else; any other value means
	// the Visitor fired the wrong number of times.
	pendingVisits int
}

// RunPayloadWalker loads the named protobuf message and recursively walks its
// fields, recording the results in the returned PayloadWalker.
func RunPayloadWalker(t *testing.T, fullName string) *PayloadWalker {
	t.Helper()

	root := mustGetProtoByName(t, fullName)
	w := &PayloadWalker{
		t:        t,
		root:     root,
		rootName: fullName,
		path:     []string{fullName},
	}
	w.walk(root.ProtoReflect())
	require.Equal(t, 0, w.pendingVisits, "walk finished with unvisited payloads:\n%s", w.Report())
	return w
}

// PayloadCount is the number of payload containers the walk reached.
func (w *PayloadWalker) PayloadCount() int { return len(w.PayloadsSeen) }

// RequirePayloadCount asserts the walk found exactly want payload containers,
// printing every payload it did find when it did not. Prefer this over
// comparing PayloadCount by hand for better error messages.
func (w *PayloadWalker) RequirePayloadCount(want int) {
	w.t.Helper()

	// Hello! Test failing? That doesn't necessarily mean you did something
	// wrong. If you have added/removed a proto field then the hard-coded number
	// of Payloads we look for in a test might have changed.
	//
	// Double check  that is the case (e.g. w.Report() includes a new proto field
	// you added).
	require.Equal(
		w.t, want, w.PayloadCount(),
		"expected %d payloads reachable from %s, found %d\n%s",
		want, w.rootName, w.PayloadCount(), w.Report())
}

// Report renders the walk as a human readable summary.
func (w *PayloadWalker) Report() string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "%s: %d messages visited, %d payloads seen\n",
		w.rootName, w.MessagesVisited, w.PayloadCount())
	sb.WriteString("payloads seen:\n")

	for i, p := range w.PayloadsSeen {
		fmt.Fprintf(&sb, "  %3d. %s\n", i+1, p)
	}

	if len(w.Skipped) > 0 {
		sb.WriteString("subtrees skipped:\n")
		for _, s := range w.Skipped {
			fmt.Fprintf(&sb, "  - %s\n", s)
		}
	}
	return sb.String()
}

// walk descends into a single message, populating one field at a time.
func (w *PayloadWalker) walk(m protoreflect.Message) {
	fullName := string(m.Descriptor().FullName())

	// Non-Temporal messages cannot contain payloads. The `Any` type is the exception,
	// it can wrap one, and the visitor is expected to recurse through it.
	if !strings.HasPrefix(fullName, temporalPrefix) && fullName != anyFullName {
		return
	}
	require.LessOrEqual(w.t, len(w.path), maxWalkDepth,
		"walk exceeded a depth of %d, which suggests a reference cycle in the protos at %s",
		maxWalkDepth, w.pathString())
	w.MessagesVisited++

	// Base case: this message is (or, for Any, can be made to hold) a payload
	// container, so stop descending and check the visitor reaches it.
	switch fullName {
	case payloadFullName, payloadsFullName:
		w.visitFromRoot(fullName)
		return

	case anyFullName:
		wrapper, ok := m.Interface().(*anypb.Any)
		require.True(w.t, ok, "expected an *anypb.Any at %s, got %T", w.pathString(), m.Interface())
		if wrapper.TypeUrl == "" {
			// Stuff the Any with an arbitrary message we know holds a payload, so
			// that reaching this payload requires recursing through the Any.
			stuffed, err := anypb.New(&update.Request{Input: &update.Input{Args: &common.Payloads{
				Payloads: []*common.Payload{{Data: []byte("orig-val")}},
			}}})
			require.NoError(w.t, err)
			proto.Merge(wrapper, stuffed)
		}
		w.visitFromRoot(fullName)
		return
	}

	fields := m.Descriptor().Fields()
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		w.walkField(m, fd)
		// Everything populated under this field has been cleared again, so any
		// payload announced while descending must have been reported back by now.
		require.Equal(w.t, 0, w.pendingVisits,
			"VisitPayloads did not reach the payload under %s.%s", w.pathString(), fd.Name())
	}
}

// walkField populates fd on m, descends into whatever it created, then clears it
// so that only one payload is ever reachable from the root at a time.
func (w *PayloadWalker) walkField(m protoreflect.Message, fd protoreflect.FieldDescriptor) {
	w.path = append(w.path, string(fd.Name()))
	defer func() { w.path = w.path[:len(w.path)-1] }()

	switch {
	case fd.IsMap():
		w.walkMapField(m, fd)
	case fd.IsList():
		w.walkListField(m, fd)
	case fd.Kind() == protoreflect.MessageKind:
		w.walkMessageField(m, fd)
	}
	// Scalar fields cannot hold payloads, so there is nothing to populate.
}

// walkMessageField handles a singular message field, including the members of a
// oneof (each oneof member is populated and cleared in turn).
func (w *PayloadWalker) walkMessageField(m protoreflect.Message, fd protoreflect.FieldDescriptor) {
	// A field of the same type as its parent (e.g. Failure.cause) would recurse
	// forever. Payloads underneath it are reachable through the parent anyway.
	if fd.Message().FullName() == m.Descriptor().FullName() {
		w.skip("self-referential field")
		return
	}

	child := m.Get(fd).Message().New()
	m.Set(fd, protoreflect.ValueOfMessage(child))
	w.walk(child)
	m.Clear(fd)
}

// walkListField handles a repeated field by appending a single element.
func (w *PayloadWalker) walkListField(m protoreflect.Message, fd protoreflect.FieldDescriptor) {
	if fd.Kind() != protoreflect.MessageKind {
		return
	}
	list := m.Mutable(fd).List()
	require.Equal(w.t, 0, list.Len(), "walker expects an empty list at %s", w.pathString())

	elem := list.NewElement()
	list.Append(elem)
	w.walk(elem.Message())
	list.Truncate(0)
}

// walkMapField handles a map field by setting a single entry.
func (w *PayloadWalker) walkMapField(m protoreflect.Message, fd protoreflect.FieldDescriptor) {
	if fd.MapValue().Kind() != protoreflect.MessageKind {
		return
	}
	mapVal := m.Mutable(fd).Map()
	require.Equal(w.t, 0, mapVal.Len(), "walker expects an empty map at %s", w.pathString())

	key := mustGetMapKey(w.t, fd.MapKey().Kind())

	// For map<_, Payload> fields, populate real payload data rather than an empty
	// message, matching how callers actually use them (e.g. Header.fields).
	entry := mapVal.NewValue()
	if fd.MapValue().Message().FullName() == payloadFullName {
		entry = protoreflect.ValueOfMessage(inputPayload().ProtoReflect())
	}
	mapVal.Set(key, entry)
	w.walk(entry.Message())
	mapVal.Clear(key)
}

// visitFromRoot records a payload container at the current path and asserts that
// VisitPayloads, run against the root message, invokes the Visitor exactly once.
//
// Exactly one payload is populated at this point, so exactly one call is
// expected: zero calls means the visitor cannot reach this field, more than one
// means it visits it repeatedly.
func (w *PayloadWalker) visitFromRoot(leafType string) {
	w.PayloadsSeen = append(w.PayloadsSeen, fmt.Sprintf("%s [%s]", w.pathString(), leafType))

	w.pendingVisits++

	ctx := context.Background()
	err := VisitPayloads(ctx, w.root, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			require.Equal(w.t, 1, w.pendingVisits,
				"VisitPayloads invoked the Visitor more than once for %s", w.pathString())
			w.pendingVisits--
			return p, nil
		},
	})
	require.NoError(w.t, err)
}

func (w *PayloadWalker) skip(reason string) {
	w.Skipped = append(w.Skipped, fmt.Sprintf("%s (%s)", w.pathString(), reason))
}

func (w *PayloadWalker) pathString() string {
	return strings.Join(w.path, ".")
}

// mustGetMapKey returns an arbitrary map key of the given kind. Panics
// if the kind is not supported.
func mustGetMapKey(t *testing.T, kind protoreflect.Kind) protoreflect.MapKey {
	switch kind {
	case protoreflect.StringKind:
		return protoreflect.ValueOfString("sample_key").MapKey()
	case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
		return protoreflect.ValueOfInt32(1).MapKey()
	case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Sfixed64Kind:
		return protoreflect.ValueOfInt64(1).MapKey()
	case protoreflect.Uint64Kind:
		return protoreflect.ValueOfUint64(1).MapKey()
	case protoreflect.BoolKind:
		return protoreflect.ValueOfBool(true).MapKey()
	}

	t.Fatalf("Map key kind %v not implemented. Please add support for it.", kind)
	return protoreflect.MapKey{}
}

// mustGetProtoByName fetches the proto message from its qualified name,
// returning its proto.Message interface.
func mustGetProtoByName(t *testing.T, fullName string) proto.Message {
	t.Helper()
	mt, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(fullName))
	require.NoError(t, err, "message %s not found in the global proto registry", fullName)
	return mt.New().Interface()
}
