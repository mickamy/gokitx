package ptr_test

import (
	"testing"

	"github.com/mickamy/gokitx/ptr"
)

func TestOf(t *testing.T) {
	t.Parallel()
	val := 42
	ptrVal := ptr.Of(val)

	if ptrVal == nil {
		t.Fatalf("Of() returned nil pointer")
	}
	if *ptrVal != val {
		t.Errorf("Of() = %v, want %v", *ptrVal, val)
	}
}

func TestUnwrap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		ptr  *int
		want int
	}{
		{
			name: "nil pointer",
			ptr:  nil,
			want: 0,
		},
		{
			name: "non-nil pointer",
			ptr:  func() *int { i := 7; return &i }(),
			want: 7,
		},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := ptr.Unwrap(tc.ptr)
			if got != tc.want {
				t.Errorf("Unwrap() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name string
		ptr  *int
		f    func(int) string
		want *string
	}{
		{
			name: "nil pointer",
			ptr:  nil,
			f: func(i int) string {
				return string(rune(i + '0'))
			},
			want: nil,
		},
		{
			name: "non-nil pointer",
			ptr:  func() *int { i := 5; return &i }(),
			f: func(i int) string {
				return string(rune(i + '0'))
			},
			want: func() *string { s := "5"; return &s }(),
		},
	}
	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := ptr.Map(tc.ptr, tc.f)
			if (got == nil) != (tc.want == nil) {
				t.Errorf("Map() got = %v, want %v", got, tc.want)
				return
			}
			if got != nil && *got != *tc.want {
				t.Errorf("Map() got = %v, want %v", *got, *tc.want)
			}
		})
	}
}
