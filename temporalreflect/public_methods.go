package temporalreflect

import (
	"reflect"
	"regexp"
)

var publicMethodRgx = regexp.MustCompile("^[A-Z]")

// PublicMethods calls the provided callback on each method declared as public on the
// specified object.
// This prevents the `mustEmbedUnimplementedFooBarBaz` method required by the GRPC v2
// gateway from polluting our tests.
func PublicMethods(obj any, cb func(reflect.Method)) {
	t := reflect.ValueOf(obj).Type()
	for i := 0; i < t.NumMethod(); i++ {
		if publicMethodRgx.MatchString(t.Method(i).Name) {
			cb(t.Method(i))
		}
	}
}
