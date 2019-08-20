package anagrams

import (
	ds "algo_design/datastructures"
	util "algo_design/util"
	"bufio"
	"os"
	"strings"
)

// Anagrams returns all anagrams of the input string using words from the line-separated word list
// found at dpath
func Anagrams(s string, dpath string) ([]string, error) {
	d, err := readDict(dpath)
	if err != nil {
		panic(err)
	}
	l := strings.ReplaceAll(s, " ", "")
	l = util.SortString(l)
	res := []string{}
	return *findAnagrams(l, &res, d, s), nil
}

func findAnagrams(l string, res *[]string, d *ds.TrieNode, s string) *[]string {
	return nil
}

func readDict(dpath string) (*ds.TrieNode, error) {
	f, err := os.Open(dpath)
	if err != nil {
		return nil, err
	}

	dsf := ds.Factory{}
	trie := dsf.TrieNode()

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		trie.Add(sc.Text())
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return trie, nil
}
