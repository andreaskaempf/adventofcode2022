// Unit tests for this Advent of Code submission

package main

import (
	"testing"
)

func TestMain(t *testing.T) {

	if 1 != 2 {
		t.Errorf("Got %d instead of 2", 1)
	}
}

// Compare two int arrays element-by-element, and report
// if they are the same (TODO: make generic)
func same(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
