package either

import (
	"fmt"
)

func Must[T any](val T, err error) T {
	if err != nil {
		panic(fmt.Errorf("must: %w", err))
	}
	return val
}

func Error[T any](_ T, err error) error {
	if err == nil {
		panic("err is nil")
	}
	return err
}

func Left[T any, U any](val T, _ U) T {
	return val
}

func Right[T any, U any](_ T, val U) U {
	return val
}

func MapLeft[T any, U any, V any](val T, _ U, f func(T) V) V {
	return f(val)
}

func MapRight[T any, U any, V any](_ T, val U, f func(U) V) V {
	return f(val)
}

func FlatMapLeft[T any, U any, V any](val T, _ U, f func(T) (V, error)) (V, error) {
	return f(val)
}

func FlatMapRight[T any, U any, V any](_ T, val U, f func(U) (V, error)) (V, error) {
	return f(val)
}
