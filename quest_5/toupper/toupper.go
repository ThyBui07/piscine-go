package piscine

func ToUpper(s string) string {
	stringToRune := []rune(s)
	newR := []rune{}
	for _, v := range stringToRune {
		if 97 <= v && v <= 122 {
			v = v - 32
			newR = append(newR, v)
		} else {
			newR = append(newR, v)
			continue
		}
	}
	return string(newR)
}
