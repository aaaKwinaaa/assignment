package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_Permutations_1(t *testing.T) {
	result := Permutations("a")
	expect := []string{"a"}

	sort.Strings(result)
	sort.Strings(expect)

	if !reflect.DeepEqual(result, expect) {
		t.Error("Result Incorrect")
	}
}

func Test_Permutations_2(t *testing.T) {
	result := Permutations("ab")
	expect := []string{"ab", "ba"}

	sort.Strings(result)
	sort.Strings(expect)

	if !reflect.DeepEqual(result, expect) {
		t.Error("Result Incorrect")
	}
}

func Test_Permutations_3(t *testing.T) {
	result := Permutations("abc")
	expect := []string{"abc", "acb", "bac", "bca", "cab", "cba"}

	sort.Strings(result)
	sort.Strings(expect)

	if !reflect.DeepEqual(result, expect) {
		t.Error("Result Incorrect")
	}
}

func Test_Permutations_4(t *testing.T) {
	result := Permutations("aabb")
	expect := []string{"aabb", "abab" , "abba" , "baab" , "baba" , "bbaa"}
	
	sort.Strings(result)
	sort.Strings(expect)

	if !reflect.DeepEqual(result, expect) {
		t.Error("Result Incorrect")
	}
}
