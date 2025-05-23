// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package enums

import (
	"fmt"
)

var (
	NexusHandlerErrorRetryBehavior_shorthandValue = map[string]int32{
		"Unspecified":  0,
		"Retryable":    1,
		"NonRetryable": 2,
	}
)

// NexusHandlerErrorRetryBehaviorFromString parses a NexusHandlerErrorRetryBehavior value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to NexusHandlerErrorRetryBehavior
func NexusHandlerErrorRetryBehaviorFromString(s string) (NexusHandlerErrorRetryBehavior, error) {
	if v, ok := NexusHandlerErrorRetryBehavior_value[s]; ok {
		return NexusHandlerErrorRetryBehavior(v), nil
	} else if v, ok := NexusHandlerErrorRetryBehavior_shorthandValue[s]; ok {
		return NexusHandlerErrorRetryBehavior(v), nil
	}
	return NexusHandlerErrorRetryBehavior(0), fmt.Errorf("%s is not a valid NexusHandlerErrorRetryBehavior", s)
}
