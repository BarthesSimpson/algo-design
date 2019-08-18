package derangements

import (
	"errors"
	"fmt"
)

/*
Use backtracking to efficiently generate all derangments of an interval {0..n}.
A derangement is a permutation P of length n such that for all i in n, P[i] != i.
*/

// AllDerangements returns a slice of slices, each representing a unique derangement
// of the numbers in the closed interval [0..n)
func AllDerangements(n int) ([][]int, error) {
	if n < 2 {
		return nil, errors.New("n must be at least 2")
	}
	res := [][]int{}
	for i := 1; i < n; i++ {
		res = append(res, []int{i})
	}
	return *backtrack(1, n, &res), nil
}

func backtrack(s, n int, res *[][]int) *[][]int {
	fmt.Printf("Iteration %d\n", s)
	fmt.Printf("%#v\n", *res)
	// once the length of our partial solutions is n, we are done
	if s == n {
		return res
	}
	// for each partial solution in res, generate all valid next steps
	next := [][]int{}
	for _, p := range *res {
		for i := 0; i < n; i++ {
			fmt.Printf("Candidate %d\n", i)
			fmt.Printf("Partial solution %#v\n", p)
			if i != s && !includes(p, i) {
				fmt.Printf("Candidate %d accepted\n", i)
				pp := append(p, i)
				next = append(next, pp)
			}
		}
	}
	return backtrack(s+1, n, &next)
}

func includes(ints []int, i int) bool {
	for _, j := range ints {
		if j == i {
			return true
		}
	}
	return false
}
