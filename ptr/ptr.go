package ptr

func Of[T any](val T) *T {
	return &val
}

func Unwrap[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}
	return *ptr
}

func Map[T any, U any](ptr *T, f func(T) U) *U {
	if ptr == nil {
		return nil
	}
	u := f(*ptr)
	return &u
}
