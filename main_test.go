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

// Test tables is a set(slice/array) of test input and output values
func TestTableSum(t *testing.T) {
	tables := []struct {
		x int
		y int
		n int
	}{
		{1, 1, 2},
		{1, 2, 3},
		{2, 2, 4},
		{5, 2, 7},
	}

	for _, table := range tables {
		total := Sum(table.x, table.y)
		if total != table.n {
			t.Errorf("Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.x, table.y, total, table.n)
		}
	}
}

/**
	go test
	go test -v
	go test -cover

	To generate HTML coverate report (cannot find some package stuff)
	go test -cover -coverprofile=c.out
	go tool cover -html=c.out -o coverage.html
**/
