package util

import "sort"

// SortString sorts a string lexographically
func SortString(s string) string {
	a := make([]byte, len(s))
	copy(a[:], s)
	sort.Slice(a, func(i int, j int) bool { return a[i] < a[j] })
	return string(a)
}
