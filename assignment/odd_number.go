package main

func OddNumber(input []int) int {
	var result int

	allNumber := make(map[int]int)

	for _, number := range input {
		allNumber[number] = allNumber[number] + 1
	}

	for number, count := range allNumber {
		if count%2 == 1 {
			result = number
			break
		}
	}

	return result
}
