package derangements

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestAllDerangements(t *testing.T) {
	g := Goblin(t)
	g.Describe("AllDerangements tests", func() {
		g.It("returns an error for n = 0", func() {
			_, err := AllDerangements(0)
			g.Assert(err != nil).IsTrue()
		})
		g.It("returns an error for n = 1", func() {
			_, err := AllDerangements(1)
			g.Assert(err != nil).IsTrue()
		})
		g.It("returns all derangements for n = 2", func() {
			res, err := AllDerangements(2)
			exp := [][]int{[]int{1, 0}}
			g.Assert(err == nil).IsTrue()
			g.Assert(exp).Equal(res)
		})
		g.It("returns all derangements for n = 3", func() {
			res, err := AllDerangements(3)
			exp := [][]int{[]int{1, 2, 0}, []int{2, 0, 1}}
			g.Assert(err == nil).IsTrue()
			g.Assert(exp).Equal(res)
		})
		g.It("returns all derangements for n = 4", func() {
			res, err := AllDerangements(4)
			exp := [][]int{
				[]int{1, 0, 3, 2}, []int{1, 2, 3, 0}, []int{1, 3, 0, 2},
				[]int{2, 0, 3, 1}, []int{2, 3, 0, 1}, []int{2, 3, 1, 0},
				[]int{3, 0, 1, 2}, []int{3, 2, 0, 1}, []int{3, 2, 1, 0},
			}
			g.Assert(err == nil).IsTrue()
			g.Assert(exp).Equal(res)
		})
	})
}
