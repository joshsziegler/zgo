package test

import (
	"reflect"
	"testing"
)

// Equals is a testing helper does a deep comparison of passed in interfaces
// and fails test if they are not equal.
func Equals(t *testing.T, actual interface{}, expected interface{}) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected and actual values do not match\n"+
			"Expected: %+v\nActual:   %+v\n", expected, actual)
	}
}
