package main

func Permutations(text string) []string {
	result := make(map[string]bool)
	permute([]rune(text), 0, len(text)-1, result)

	allPossible := []string{}
	for key := range result {
		allPossible = append(allPossible, key)
	}
	return allPossible
}

func permute(text []rune, startIndex, lastIndex int, result map[string]bool) {
	if startIndex == lastIndex {
		result[string(text)] = true
	} else {
		for i := startIndex; i <= lastIndex; i++ {
			text[startIndex], text[i] = text[i], text[startIndex]
			permute(text, startIndex+1, lastIndex, result)
			text[startIndex], text[i] = text[i], text[startIndex]
		}
	}

}
