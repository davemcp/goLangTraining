package word

import (
	"testing"
)

const test_string = "Lorem ipsum another word"
func TestCount(t *testing.T) {
	c := Count(test_string)
	if c != 4 {
		t.Error("Expected", 4, "received", c)
	}
}
func TestUseCount(t *testing.T) {
	c := UseCount(test_string)
	if c["ipsum"] != 1 {
		t.Error("Expected", 1, "received", c)
	}
	if len(c) != 4 {
		t.Error("Expected", 4, "received", len(c))
	}
}