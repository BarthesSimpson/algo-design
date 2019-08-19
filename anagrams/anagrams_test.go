package anagrams

import (
	"testing"

	. "github.com/franela/goblin"
)

func TestAllDerangements(t *testing.T) {
	g := Goblin(t)
	g.Describe("Anagrams tests", func() {
		g.It("converts \"many voted bush retired\" to \"tuesday november third\"", func() {
			res, err := Anagrams("many voted bush retired", "./test_dictionary.txt")
			exp := []string{"tuesday november third"}
			g.Assert(err != nil).IsTrue()
			g.Assert(exp).Equal(res)
		})
		g.It("converts \"tuesday november third\" to \"many voted bush retired\"", func() {
			res, err := Anagrams("tuesday november third", "./test_dictionary.txt")
			exp := []string{"many voted bush retired"}
			g.Assert(err != nil).IsTrue()
			g.Assert(exp).Equal(res)
		})
	})
}