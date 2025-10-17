package operator_test

import (
	"testing"

	"github.com/mickamy/gokitx/operator"
)

func TestTernary(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		cond     bool
		a, b     int
		expected int
	}{
		{true, 1, 2, 1},
		{false, 1, 2, 2},
		{true, -5, 5, -5},
		{false, -5, 5, 5},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run("", func(t *testing.T) {
			t.Parallel()
			got := operator.Ternary(tc.cond, tc.a, tc.b)
			if got != tc.expected {
				t.Errorf("Ternary(%v, %d, %d) = %d; want %d", tc.cond, tc.a, tc.b, got, tc.expected)
			}
		})
	}
}

func TestTernaryFunc(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		cond     bool
		a, b     int
		expected int
	}{
		{true, 1, 2, 1},
		{false, 1, 2, 2},
		{true, -5, 5, -5},
		{false, -5, 5, 5},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run("", func(t *testing.T) {
			t.Parallel()
			got := operator.TernaryFunc(tc.cond, func() int { return tc.a }, func() int { return tc.b })
			if got != tc.expected {
				t.Errorf("TernaryFunc(%v, %d, %d) = %d; want %d", tc.cond, tc.a, tc.b, got, tc.expected)
			}
		})
	}
}
