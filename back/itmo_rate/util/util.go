package util

func Map[T, U any](data []T, f func(T) U) []U {
	ret := make([]U, len(data))
	for i, d := range data {
		ret[i] = f(d)
	}
	return ret
}

func Filter[T any](data []T, f func(T) bool) []T {
	ret := []T{}
	for _, d := range data {
		if f(d) {
			ret = append(ret, d)
		}
	}
	return ret
}
