// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
