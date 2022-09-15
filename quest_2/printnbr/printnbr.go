package piscine

import (
	"github.com/01-edu/z01"
)

func printDigit1(n int) []int {
	var s []int
	var r int
	for n != 0 {

		r = n % 10
		s = append(s, r)

		n = n / 10
	}
	return s
}

func PrintNbr(n int) {
	var neg bool

	if n < 0 {
		neg = true
	}

	if n == 0 {
		z01.PrintRune('0')
	}
	s := printDigit1(n)
	if neg == true {
		z01.PrintRune('-')
	}
	for i, v := range s {
		v = s[len(s)-i-1]
		if neg == true {
			v = v * -1
		}
		z01.PrintRune(rune(v + 48))
	}
}
