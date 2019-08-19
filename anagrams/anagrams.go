package anagrams

import (
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

func findAnagrams(l string, res *[]string, d map[string]bool, s string) *[]string {
	return nil
}

func readDict(dpath string) (map[string]bool, error) {
	f, err := os.Open(dpath)
	if err != nil {
		return nil, err
	}

	d := map[string]bool{}

	sc := bufio.NewScanner(f)
	for sc.Scan() {
		d[sc.Text()] = true
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return d, nil
}
