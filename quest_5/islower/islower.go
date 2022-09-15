package piscine

func IsLower(s string) bool {
	rune := []rune(s)
	res := true
	for _, v := range rune {
		if v < 'a' || v > 'z' {
			res = false
		}
	}
	return res
}
