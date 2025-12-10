package utils

func CompareEqualLenIntSlice(a, b []int) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
