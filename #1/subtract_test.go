package math

import (
	"testing"
	"testing/quick"
)

func TestSubtract(t *testing.T) {
	testCases := []struct {
		a, b   int
		expect int
	}{
		{1, -1, 2},
		{0, -1, 1},
		{1, -4, 5},
	}
	for _, tc := range testCases {
		if result := Subtract(tc.a, tc.b); result != tc.expect {
			t.Errorf("Subtract() = %v, want %v", result, tc.expect)
		}
	}
}

func TestSubtractPropertyFour(t *testing.T) {
	property := func(a, b int) bool {
		if b > 0 {
			return true
		}
		return Subtract(a, b) > a
	}
	if err := quick.Check(property, nil); err != nil {
		t.Error(err)
	}
}
