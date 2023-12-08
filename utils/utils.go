package utils

func Pop[T ~string | ~int | float64](s []T) (T, []T) {
	return s[len(s)-1], s[:len(s)-1]
}

func Push[T ~string | ~int | float64](s []T, v T) []T {
	return append(s, v)
}
