package piscine

func StrRev(s string) string {
	var word string
	for index := range s {
		wordASC := s[len(s)-index-1]
		word += string(wordASC)
	}
	return word
}
