package piscine

func IsUpper(s string) bool {
	rune := []rune(s)
	res := true
	for _, v := range rune {
		if v < 'A' || v > 'Z' {
			res = false
		}
	}
	return res
}
