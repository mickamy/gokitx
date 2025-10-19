package either_test

import (
	"errors"
	"testing"

	"github.com/mickamy/gokitx/either"
)

func TestMust(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		val      int
		err      error
		expected int
		panic    bool
	}{
		{42, nil, 42, false},
		{0, nil, 0, false},
		{0, errors.New("some error"), 0, true},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run("", func(t *testing.T) {
			t.Parallel()
			defer func() {
				r := recover()
				if (r != nil) != tc.panic {
					t.Errorf("Must(%d, %v) panic = %v; want panic = %v", tc.val, tc.err, r != nil, tc.panic)
				}
			}()
			got := either.Must(tc.val, tc.err)
			if got != tc.expected {
				t.Errorf("Must(%d, %v) = %d; want %d", tc.val, tc.err, got, tc.expected)
			}
		})
	}
}

func TestError(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		val      int
		err      error
		expected error
		panic    bool
	}{
		{42, errors.New("some error"), errors.New("some error"), false},
		{0, errors.New("another error"), errors.New("another error"), false},
		{0, nil, nil, true},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run("", func(t *testing.T) {
			t.Parallel()
			defer func() {
				r := recover()
				if (r != nil) != tc.panic {
					t.Errorf("Error(%d, %v) panic = %v; want panic = %v", tc.val, tc.err, r != nil, tc.panic)
				}
			}()
			got := either.Error(tc.val, tc.err)
			if got == nil && tc.expected == nil {
				return
			}
			if got == nil || tc.expected == nil || got.Error() != tc.expected.Error() {
				t.Errorf("Error(%d, %v) = %v; want %v", tc.val, tc.err, got, tc.expected)
			}
		})
	}
}

func TestLeft(t *testing.T) {
	t.Parallel()

	valLeft := 10
	valRight := "right value"

	got := either.Left(valLeft, valRight)

	if got != valLeft {
		t.Errorf("Left(%d, %q) = %d; want %d", valLeft, valRight, got, valLeft)
	}
}

func TestRight(t *testing.T) {
	t.Parallel()

	valLeft := 10
	valRight := "right value"

	got := either.Right(valLeft, valRight)

	if got != valRight {
		t.Errorf("Right(%d, %q) = %q; want %q", valLeft, valRight, got, valRight)
	}
}

func TestRecover(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name      string
		fn        func() int
		expectErr bool
		expected  int
	}{
		{
			name: "no panic",
			fn: func() int {
				return 42
			},
			expectErr: false,
			expected:  42,
		},
		{
			name: "with panic",
			fn: func() int {
				panic("something went wrong")
			},
			expectErr: true,
			expected:  0,
		},
		{
			name: "panic with error",
			fn: func() int {
				panic(errors.New("an error occurred"))
			},
			expectErr: true,
			expected:  0,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got, err := either.Recover(tc.fn)
			if tc.expectErr {
				if err == nil {
					t.Errorf("Recover() expected error, got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Recover() unexpected error: %v", err)
				}
			}
			if got != tc.expected {
				t.Errorf("Recover() = %d; want %d", got, tc.expected)
			}
		})
	}
}

func TestMapLeft(t *testing.T) {
	t.Parallel()

	valLeft := 5
	valRight := "unused"

	got := either.MapLeft(valLeft, valRight, func(v int) int {
		return v * 2
	})

	expected := 10
	if got != expected {
		t.Errorf("MapLeft(%d, %q, f) = %d; want %d", valLeft, valRight, got, expected)
	}
}

func TestMapRight(t *testing.T) {
	t.Parallel()

	valLeft := 0
	valRight := "hello"

	got := either.MapRight(valLeft, valRight, func(v string) string {
		return v + " world"
	})

	expected := "hello world"
	if got != expected {
		t.Errorf("MapRight(%d, %q, f) = %q; want %q", valLeft, valRight, got, expected)
	}
}

func TestFlatMapLeft(t *testing.T) {
	t.Parallel()

	valLeft := 3
	valRight := "unused"

	got, err := either.FlatMapLeft(valLeft, valRight, func(v int) (int, error) {
		if v < 0 {
			return 0, errors.New("negative value")
		}
		return v * 3, nil
	})

	if err != nil {
		t.Errorf("FlatMapLeft(%d, %q, f) returned error: %v", valLeft, valRight, err)
		return
	}

	expected := 9
	if got != expected {
		t.Errorf("FlatMapLeft(%d, %q, f) = %d; want %d", valLeft, valRight, got, expected)
	}
}

func TestFlatMapRight(t *testing.T) {
	t.Parallel()

	valLeft := 0
	valRight := "go"

	got, err := either.FlatMapRight(valLeft, valRight, func(v string) (string, error) {
		if v == "" {
			return "", errors.New("empty string")
		}
		return v + "lang", nil
	})

	if err != nil {
		t.Errorf("FlatMapRight(%d, %q, f) returned error: %v", valLeft, valRight, err)
		return
	}

	expected := "golang"
	if got != expected {
		t.Errorf("FlatMapRight(%d, %q, f) = %q; want %q", valLeft, valRight, got, expected)
	}
}
