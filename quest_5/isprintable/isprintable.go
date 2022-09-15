package piscine

func IsPrintable(s string) bool {
	rune := []rune(s)
	res := true
	for _, v := range rune {
		if v < 32 || v > 127 {
			res = false
		}
	}
	return res
}
