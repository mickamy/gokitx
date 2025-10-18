package slices

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func FlatMap[T, U any](ts []T, f func(T) []U) []U {
	var result []U
	for _, t := range ts {
		result = append(result, f(t)...)
	}
	return result
}

func GroupBy[T any, KEY comparable](items []T, getProperty func(T) KEY) map[KEY][]T {
	grouped := make(map[KEY][]T)

	for _, item := range items {
		key := getProperty(item)
		grouped[key] = append(grouped[key], item)
	}

	return grouped
}

func Find[T any](ts []T, f func(T) bool) (T, bool) {
	for _, t := range ts {
		if f(t) {
			return t, true
		}
	}
	var zero T
	return zero, false
}

func Filter[T any](ts []T, f func(T) bool) []T {
	var result []T
	for _, t := range ts {
		if f(t) {
			result = append(result, t)
		}
	}
	return result
}

func Unique[T comparable](ts []T) []T {
	seen := make(map[T]struct{})
	var result []T
	for _, t := range ts {
		if _, ok := seen[t]; !ok {
			seen[t] = struct{}{}
			result = append(result, t)
		}
	}
	return result
}
