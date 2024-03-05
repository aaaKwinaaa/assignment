package main

import "testing"

func Test_Smiley_1(t *testing.T) {
	input := []string{":)", ";(", ";}", ":-D"}
	result := Smiley(input)
	if result != 2 {
		t.Error("Result Incorrect")
	}
}

func Test_Smiley_2(t *testing.T) {
	input := []string{";D", ":-(", ":-)", ";~)"}
	result := Smiley(input)
	if result != 3 {
		t.Error("Result Incorrect")
	}
}

func Test_Smiley_3(t *testing.T) {
	input := []string{";]", ":[", ";*", ":$", ";-D"}
	result := Smiley(input)
	if result != 1 {
		t.Error("Result Incorrect")
	}
}
