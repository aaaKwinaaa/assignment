package main

import "testing"

func Test_OodNumber_1(t *testing.T) {
	input := []int{7}
	result := OddNumber(input)
	if result != 7 {
		t.Error("Result Incorrect")
	}
}

func Test_OodNumber_2(t *testing.T) {
	input := []int{0}
	result := OddNumber(input)
	if result != 0 {
		t.Error("Result Incorrect")
	}
}

func Test_OodNumber_3(t *testing.T) {
	input := []int{1,1,2}
	result := OddNumber(input)
	if result != 2 {
		t.Error("Result Incorrect")
	}
}

func Test_OodNumber_4(t *testing.T) {
	input := []int{0,1,0,1,0}
	result := OddNumber(input)
	if result != 0 {
		t.Error("Result Incorrect")
	}
}

func Test_OodNumber_5(t *testing.T) {
	input := []int{1,2,2,3,3,3,4,3,3,3,2,2,1}
	result := OddNumber(input)
	if result != 4 {
		t.Error("Result Incorrect")
	}
}


