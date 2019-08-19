package util

import "testing"

func TestSort(t *testing.T) {
	exp := "abcdefg"
	res := SortString("gfedcba")
	if exp != res {
		t.Fail()
	}
}

func TestSortEmpty(t *testing.T) {
	exp := ""
	res := SortString("")
	if exp != res {
		t.Fail()
	}
}

func TestSortNumeric(t *testing.T) {
	exp := "123abc"
	res := SortString("321cba")
	if exp != res {
		t.Fail()
	}
}
