package piscine

func ToLower(s string) string {
	stringToRune := []rune(s)
	newR := []rune{}
	for _, v := range stringToRune {
		if 65 <= v && v <= 90 {
			v = v + 32
			newR = append(newR, v)
		} else {
			newR = append(newR, v)
			continue
		}
	}
	return string(newR)
}
