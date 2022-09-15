package piscine

func Index(s string, toFind string) int {
	if len(s) == 0 || len(toFind) == 0 {
		return 0
	}
	toFindRune := []rune(toFind)
	string1 := []rune(s)
	res := 0
	for i, v := range string1 {
		if v != toFindRune[0] {
			res = -1
			continue
		} else {
			res = i
			break
		}
	}
	return res
}
