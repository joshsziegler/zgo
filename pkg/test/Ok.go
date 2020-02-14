package test

import "testing"

// Ok is a testing helper that will fail the test if err does not equal nil.
func Ok(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Testing Failed: %v", err)
	}
}
