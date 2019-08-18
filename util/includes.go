package util

// Includes returns true if and only if i is in ints
func Includes(ints []int, i int) bool {
	for _, j := range ints {
		if j == i {
			return true
		}
	}
	return false
}
