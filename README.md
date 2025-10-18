# gokitx

Utility helpers for Go 1.23+ that cover common patterns such as ternary evaluation, pointer helpers, and small tuple structs. The library is intentionally lightweight so you can drop in specific pieces without adopting a large framework.

## Modules

- `operator`: generic `Ternary` and `TernaryFunc` helpers that return eager or lazily computed values based on a boolean condition.
- `slices`: functional helpers like `Map`, `FlatMap`, `GroupBy`, `Find`, `Filter`, and `Unique` for working with collections.
- `ptr`: convenience functions for working with pointers, including `Of`, `Unwrap`, and `Map`.
- `tuple`: generic structs for 2–10 element tuples when you need to group multiple return values.

## Installation

```shell
go get github.com/mickamy/gokitx
```

## Usage

```go
package main

import (
	"fmt"

	"github.com/mickamy/gokitx/operator"
	"github.com/mickamy/gokitx/ptr"
	"github.com/mickamy/gokitx/slices"
	"github.com/mickamy/gokitx/tuple"
)

func main() {
	numbers := []int{1, 2, 2, 3}
	unique := slices.Unique(numbers)
	labels := slices.Map(unique, func(v int) string {
		return fmt.Sprintf("value: %d", v)
	})

	selected, ok := slices.Find(unique, func(v int) bool {
		return v > 1
	})

	n := operator.Ternary(ok, selected, 0)

	p := ptr.Map(ptr.Of(n), func(v int) string {
		return fmt.Sprintf("value: %d", v)
	})

	t := tuple.Triple[int, string, []string]{
		First:  n,
		Second: ptr.Unwrap(p),
		Third:  labels,
	}

	fmt.Println(t)
}
```

## License

[MIT](./LICENSE)
