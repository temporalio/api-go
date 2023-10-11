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

package enums

import (
	"fmt"
)

var (
	NamespaceState_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Registered":  1,
		"Deprecated":  2,
		"Deleted":     3,
	}
	NamespaceState_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Registered",
		2: "Deprecated",
		3: "Deleted",
	}
)

// NamespaceStateFromString parses a NamespaceState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to NamespaceState
func NamespaceStateFromString(s string) (NamespaceState, error) {
	if v, ok := NamespaceState_value[s]; ok {
		return NamespaceState(v), nil
	} else if v, ok := NamespaceState_shorthandValue[s]; ok {
		return NamespaceState(v), nil
	}
	return NamespaceState(0), fmt.Errorf("%s is not a valid NamespaceState", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	NamespaceState(-1).Shorthand() // will return "", false
func (e NamespaceState) Shorthand() (string, bool) {
	if s, ok := NamespaceState_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	ArchivalState_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Disabled":    1,
		"Enabled":     2,
	}
	ArchivalState_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Disabled",
		2: "Enabled",
	}
)

// ArchivalStateFromString parses a ArchivalState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ArchivalState
func ArchivalStateFromString(s string) (ArchivalState, error) {
	if v, ok := ArchivalState_value[s]; ok {
		return ArchivalState(v), nil
	} else if v, ok := ArchivalState_shorthandValue[s]; ok {
		return ArchivalState(v), nil
	}
	return ArchivalState(0), fmt.Errorf("%s is not a valid ArchivalState", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	ArchivalState(-1).Shorthand() // will return "", false
func (e ArchivalState) Shorthand() (string, bool) {
	if s, ok := ArchivalState_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	ReplicationState_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Normal":      1,
		"Handover":    2,
	}
	ReplicationState_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Normal",
		2: "Handover",
	}
)

// ReplicationStateFromString parses a ReplicationState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ReplicationState
func ReplicationStateFromString(s string) (ReplicationState, error) {
	if v, ok := ReplicationState_value[s]; ok {
		return ReplicationState(v), nil
	} else if v, ok := ReplicationState_shorthandValue[s]; ok {
		return ReplicationState(v), nil
	}
	return ReplicationState(0), fmt.Errorf("%s is not a valid ReplicationState", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	ReplicationState(-1).Shorthand() // will return "", false
func (e ReplicationState) Shorthand() (string, bool) {
	if s, ok := ReplicationState_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}
