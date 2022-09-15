package piscine

func IsNumeric(s string) bool {
	rune := []rune(s)
	res := true
	for _, v := range rune {
		if v < '0' || v > '9' {
			res = false
		}
	}
	return res
}
