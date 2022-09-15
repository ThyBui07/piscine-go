package piscine

func LastRune(s string) rune {
	stringToRune := []rune(s)
	return stringToRune[len(stringToRune)-1]
}
