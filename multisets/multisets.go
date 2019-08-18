package multisets

import (
	util "algo_design/util"
)

/*
Use backtracking to generate all permutations of the given multiset.
A multiset is a set that can multiple instances of any member.
*/

// Permutations returns a slice of slices, each representing a unique permutation
// of the multiset ms
func Permutations(ms []int) ([][]int, error) {
	if len(ms) < 2 {
		return [][]int{ms}, nil
	}
	res := [][]int{}
	seen := map[int]bool{}
	for _, i := range ms {
		if _, used := seen[i]; !used {
			res = append(res, []int{i})
			seen[i] = true
		}
	}
	cnt := util.Count(ms)
	return *backtrack(1, &ms, &res, cnt), nil
}

func backtrack(s int, ms *[]int, res *[][]int, cnt map[int]int) *[][]int {
	mss := *ms
	// if our partial solution length is the length of the multiset, we are done
	if s == len(mss) {
		return res
	}
	next := [][]int{}
	// for each partial solution, construct all valid next steps
	for _, p := range *res {
		seen := map[int]bool{}
		for i := 0; i < len(mss); i++ {
			c := mss[i]
			if _, used := seen[c]; used {
				continue
			}
			if util.Count(p)[c] == cnt[c] {
				continue
			}
			pp := append(p, c)
			next = append(next, pp)
			seen[c] = true
		}
	}
	// recurse using the new partial solution set
	return backtrack(s+1, ms, &next, cnt)
}
