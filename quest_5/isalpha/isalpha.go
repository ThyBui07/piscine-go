package piscine

func IsAlpha(s string) bool {
	rune := []rune(s)
	res := true
	for _, v := range rune {
		if (v < 'a' || v > 'z') && (v < 'A' || v > 'Z') && (v < '0' || v > '9') {
			res = false
		}
	}
	return res
}
