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

const protoGeneratedString = `package persistence

import (
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

type BuildId_State int32

const (
	BuildId_STATE_UNSPECIFIED BuildId_State = 0
	BuildId_STATE_ACTIVE      BuildId_State = 1
	BuildId_STATE_DELETED     BuildId_State = 2
)

func (x BuildId_State) Ignored() bool {
	return true
}

func (x BuildId_State) String() string {
	return protoimpl.X.MessageStringOf(x)
}

`
const stringSwitchStmt = `package persistence

import (
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	"strconv"
)

type BuildId_State int32

const (
	BuildId_STATE_UNSPECIFIED BuildId_State = 0
	BuildId_STATE_ACTIVE      BuildId_State = 1
	BuildId_STATE_DELETED     BuildId_State = 2
)

func (x BuildId_State) Ignored() bool {
	return true
}

func (x BuildId_State) String() string {
	switch x {
	case BuildId_STATE_UNSPECIFIED:
		return "BuildIdStateUnspecified"
	case BuildId_STATE_ACTIVE:
		return "BuildIdStateActive"
	case BuildId_STATE_DELETED:
		return "BuildIdStateDeleted"
	default:
		return strconv.Itoa(int(x))
	}

}
`

func TestRewriteStringMethods(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", protoGeneratedString, parser.ParseComments)
	if err != nil {
		t.Errorf("Failed to parse code: %s", err)
		t.FailNow()
	}

	cr := protogen.NewStringRewriter()
	cr.Process(f)

	var b strings.Builder
	if err := format.Node(&b, fset, f); err != nil {
		t.Errorf("Failed to format AST: %s", err)
		t.FailNow()
	}
	require.Equal(t, stringSwitchStmt, b.String())
}
