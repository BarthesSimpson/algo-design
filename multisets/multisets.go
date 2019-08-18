package multisets

import (
	"errors"
	"fmt"
)

/*
Use backtracking to generate all permutations of the given multiset.
A multiset is a set that can multiple instances of any member.
*/

// Permutations returns a slice of slices, each representing a unique permutation
// of the multiset ms
func Permutations(ms [][]int) ([][]int, error) {
	if len(ms) < 2 {
		return ms, nil
	}
	return nil, errors.New("Not yet implemented")
}

