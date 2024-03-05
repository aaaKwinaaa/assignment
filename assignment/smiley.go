package main

func Smiley(smileys []string) int {
	count := 0

	for _, smiley := range smileys {
		if len(smiley) == 2 {
			if (smiley[0] == ';' || smiley[0] == ':') && (smiley[1] == ')' || smiley[1] == 'D') {
				count++
			}
		} else { // THAT MEAN 3
			if (smiley[0] == ';' || smiley[0] == ':') && (smiley[1] == '-' || smiley[1] == '~') && (smiley[2] == ')' || smiley[2] == 'D') {
				count++
			}
		}
	}
	return count

}
