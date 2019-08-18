package util

// Count returns a map of i to frequency of i in ints for each member i of ints
func Count(ints []int) map[int]int {
	res := map[int]int{}
	for _, i := range ints {
		res[i]++
	}
	return res
}