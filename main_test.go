package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		// You can also use t.Error or t.Fail to indicate failure
		// t.Log can be used to provide non-failing debug info
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
