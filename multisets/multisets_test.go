package multisets

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestPermutations(t *testing.T) {
	g := Goblin(t)
	g.Describe("Permutations tests", func() {
		g.It("works for the empty set", func() {
			res, err := Permutations([]int{})
			exp := [][]int{[]int{}}
			g.Assert(err == nil).IsTrue()
			g.Assert(exp).Equal(res)
		})
		g.It("works for a set with one member", func() {
			res, err := Permutations([]int{1})
			exp := [][]int{[]int{1}}
			g.Assert(err == nil).IsTrue()
			g.Assert(exp).Equal(res)
		})
		g.It("returns a single result if all members in ms are identical", func() {
			res, err := Permutations([]int{1, 1, 1})
			exp := [][]int{[]int{1, 1, 1}}
			g.Assert(err == nil).IsTrue()
			g.Assert(exp).Equal(res)
		})
		g.It("works for {1, 2, 3}", func() {
			res, err := Permutations([]int{1, 2, 3})
			exp := [][]int{
				[]int{1, 2, 3}, []int{1, 3, 2},
				[]int{2, 1, 3}, []int{2, 3, 1},
				[]int{3, 1, 2}, []int{3, 2, 1},
			}
			g.Assert(err == nil).IsTrue()
			g.Assert(exp).Equal(res)
		})
		g.It("works for {1, 1, 2}", func() {
			res, err := Permutations([]int{1, 1, 2})
			exp := [][]int{
				[]int{1, 1, 2}, []int{1, 2, 1}, []int{2, 1, 1},
			}
			g.Assert(err == nil).IsTrue()
			g.Assert(exp).Equal(res)
		})
		g.It("works for {1, 1, 1, 2}", func() {
			res, err := Permutations([]int{1, 1, 1, 2})
			exp := [][]int{
				[]int{1, 1, 1, 2}, []int{1, 1, 2, 1},
				[]int{1, 2, 1, 1}, []int{2, 1, 1, 1},
			}
			g.Assert(err == nil).IsTrue()
			g.Assert(exp).Equal(res)
		})
		g.It("works for {1, 1, 2, 2}", func() {
			res, err := Permutations([]int{1, 1, 2, 2})
			exp := [][]int{
				[]int{1, 1, 2, 2}, []int{1, 2, 1, 2},
				[]int{1, 2, 2, 1}, []int{2, 1, 1, 2},
				[]int{2, 1, 2, 1}, []int{2, 2, 1, 1},
			}
			g.Assert(err == nil).IsTrue()
			g.Assert(exp).Equal(res)
		})
	})
}
