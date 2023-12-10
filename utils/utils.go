package utils

func Pop[T ~string | ~int | float64](s *[]T) T {
	if len(*s) == 0 {
		var zero T
		return zero
	}
	poppedValue := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return poppedValue
}

func Push[T ~string | ~int | float64](s *[]T, v T) {
	*s = append(*s, v)
}
