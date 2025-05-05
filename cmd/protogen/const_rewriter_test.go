package main_test

import (
	"go/format"
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	protogen "go.temporal.io/api/cmd/protogen"
)

const given = `package persistence

type BuildId_State int32

const (
	BuildId_STATE_UNSPECIFIED BuildId_State = 0
	BuildId_STATE_ACTIVE      BuildId_State = 1
	BuildId_STATE_DELETED     BuildId_State = 2
)

// Enum value maps for BuildId_State.
var (
	BuildId_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_ACTIVE",
		2: "STATE_DELETED",
	}
	BuildId_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STATE_ACTIVE":      1,
		"STATE_DELETED":     2,
	}
)

// BuildId is an identifier with a timestamped status used to identify workers for task queue versioning purposes.
type BuildId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

    State BuildId_State
}

func (x *BuildId) GetState() BuildId_State {
	if x != nil {
		return x.State
	}
	return BuildId_STATE_UNSPECIFIED
}

type EncodingType int32

const (
	ENCODING_TYPE_UNSPECIFIED EncodingType = 0
	ENCODING_TYPE_PROTO3      EncodingType = 1
	ENCODING_TYPE_JSON        EncodingType = 2
)
`

const expected = `package persistence

type BuildId_State int32

const (
	STATE_UNSPECIFIED BuildId_State = 0
	STATE_ACTIVE      BuildId_State = 1
	STATE_DELETED     BuildId_State = 2
)

// Enum value maps for BuildId_State.
var (
	BuildId_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_ACTIVE",
		2: "STATE_DELETED",
	}
	BuildId_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STATE_ACTIVE":      1,
		"STATE_DELETED":     2,
	}
)

// BuildId is an identifier with a timestamped status used to identify workers for task queue versioning purposes.
type BuildId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State BuildId_State
}

func (x *BuildId) GetState() BuildId_State {
	if x != nil {
		return x.State
	}
	return STATE_UNSPECIFIED
}

type EncodingType int32

const (
	ENCODING_TYPE_UNSPECIFIED EncodingType = 0
	ENCODING_TYPE_PROTO3      EncodingType = 1
	ENCODING_TYPE_JSON        EncodingType = 2
)
`

func TestRewriteConstEnums(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", given, parser.ParseComments)
	if err != nil {
		t.Errorf("Failed to parse code: %s", err)
		t.FailNow()
	}

	cr := protogen.NewConstRewriter(map[string]string{
		// Add a single special case. EncodingType should be discovered
		// by the tool itself
		"BuildId_State": "BuildId",
	})
	cr.Process(f)

	var b strings.Builder
	if err := format.Node(&b, fset, f); err != nil {
		t.Errorf("Failed to format AST: %s", err)
		t.FailNow()
	}
	require.Equal(t, expected, b.String())
}
