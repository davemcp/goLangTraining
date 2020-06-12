package main

import "testing"

func TestSqrRoot(t *testing.T) {
	v, _ := SqrRoot(2)
	if v != 1.4142135623730951 {
		t.Error("Expected 1.4142135623730951, got ", v)
	}
}
