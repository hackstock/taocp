package algos

import (
	"testing"
)

func TestGCD(t *testing.T) {
	testCases := []struct {
		tag string
		m   int
		n   int
		ans int
	}{
		{
			tag: "gcd of 15 and 5",
			m:   15,
			n:   5,
			ans: 5,
		},
		{
			tag: "gcd of 10 and 125",
			m:   10,
			n:   125,
			ans: 5,
		},
		{
			tag: "gcd of 56 and 42",
			m:   56,
			n:   42,
			ans: 14,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.tag, func(t *testing.T) {
			expected := tc.ans
			got := gcd(tc.m, tc.n)

			if got != expected {
				t.Fatalf("expected %d, got %d", expected, got)
			}
		})
	}
}
