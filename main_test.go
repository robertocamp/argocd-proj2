package main 

import (
	"fmt"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},	}

	for _, test := range tests {
		testname := fmt.Sprintf("%d,%d", test.a, test.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(test.a, test.b)
			if ans != test.want {
				t.Errorf("got %d, want %d", ans, test.want)
			}
		})
		if ans := IntMin(test.a, test.b); ans != test.want {
			t.Errorf("IntMin(%d, %d) = %d; want %d", test.a, test.b, ans, test.want)
		}
	}
}


