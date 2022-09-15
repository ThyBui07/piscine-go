package piscine

func NRune(s string, n int) rune {
	stringToRune := []rune(s)
	if n <= 0 || n > len(s) {
		return 0
	}
	return stringToRune[n-1]
}
