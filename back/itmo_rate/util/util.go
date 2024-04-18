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

func Mean(data []float32) float32 {
	if len(data) == 0 {
		return 0
	}
	sum := float32(0)
	for _, el := range data {
		sum += el
	}
	return sum / float32(len(data))
}
