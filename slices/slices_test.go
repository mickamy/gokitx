package slices_test

import (
	"strconv"
	"testing"

	"github.com/mickamy/gokitx/slices"
)

func TestMap(t *testing.T) {
	t.Parallel()

	// arrange
	intSlice := []int{1, 2, 3}

	// act
	got := slices.Map(intSlice, func(v int) string {
		return strconv.Itoa(v)
	})

	// assert
	want := []string{"1", "2", "3"}
	if len(got) != len(want) {
		t.Fatalf("got length %d, want length %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestFlatMap(t *testing.T) {
	t.Parallel()

	// arrange
	intSlice := []int{1, 2, 3}

	// act
	got := slices.FlatMap(intSlice, func(v int) []string {
		return []string{strconv.Itoa(v), strconv.Itoa(v * 10)}
	})

	// assert
	want := []string{"1", "10", "2", "20", "3", "30"}
	if len(got) != len(want) {
		t.Fatalf("got length %d, want length %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestGroupBy(t *testing.T) {
	t.Parallel()

	// arrange
	strSlice := []string{"apple", "banana", "apricot", "blueberry", "cherry"}

	// act
	got := slices.GroupBy(strSlice, func(s string) string {
		return string(s[0]) // Group by first letter
	})

	// assert
	want := map[string][]string{
		"a": {"apple", "apricot"},
		"b": {"banana", "blueberry"},
		"c": {"cherry"},
	}
	if len(got) != len(want) {
		t.Fatalf("got length %d, want length %d", len(got), len(want))
	}
	for k, v := range want {
		gotValues, ok := got[k]
		if !ok {
			t.Errorf("missing key %q in got", k)
			continue
		}
		if len(gotValues) != len(v) {
			t.Errorf("for key %q, got length %d, want length %d", k, len(gotValues), len(v))
			continue
		}
		for i := range v {
			if gotValues[i] != v[i] {
				t.Errorf("for key %q, got %v, want %v", k, gotValues, v)
			}
		}
	}
}

func TestFind(t *testing.T) {
	t.Parallel()

	// arrange
	intSlice := []int{1, 2, 3, 4, 5}

	// act
	got, found := slices.Find(intSlice, func(v int) bool {
		return v%2 == 0 // Find first even number
	})

	// assert
	if !found {
		t.Fatalf("expected to find an even number, but did not")
	}
	want := 2
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestFilter(t *testing.T) {
	t.Parallel()

	// arrange
	intSlice := []int{1, 2, 3, 4, 5}

	// act
	got := slices.Filter(intSlice, func(v int) bool {
		return v%2 != 0 // Filter odd numbers
	})

	// assert
	want := []int{1, 3, 5}
	if len(got) != len(want) {
		t.Fatalf("got length %d, want length %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestUnique(t *testing.T) {
	t.Parallel()

	// arrange
	intSlice := []int{1, 2, 2, 3, 3, 3}

	// act
	got := slices.Unique(intSlice)

	// assert
	want := []int{1, 2, 3}
	if len(got) != len(want) {
		t.Fatalf("got length %d, want length %d", len(got), len(want))
	}
	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %v, want %v", got, want)
		}
	}
}

func TestAll(t *testing.T) {
	t.Parallel()

	tcs := []struct {
		name     string
		intSlice []int
		pred     func(int) bool
		want     bool
	}{
		{
			name:     "all elements satisfy predicate",
			intSlice: []int{2, 4, 6},
			pred: func(v int) bool {
				return v%2 == 0
			},
			want: true,
		},
		{
			name:     "at least one element fails predicate",
			intSlice: []int{2, 3, 4},
			pred: func(v int) bool {
				return v%2 == 0
			},
			want: false,
		},
		{
			name:     "empty slice returns true",
			intSlice: []int{},
			pred: func(v int) bool {
				return v%2 == 0
			},
			want: true,
		},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := slices.All(tc.intSlice, tc.pred)

			if got != tc.want {
				t.Errorf("All() = %v, want %v", got, tc.want)
			}
		})
	}
}
