package main

import (
	"github.com/01-edu/z01"
)

func main() {
	counter := 0
	for i := 'a'; i <= 'z'; i++ {
		if counter%2 != 0 {
			z01.PrintRune(i - 32)
		} else {
			z01.PrintRune(i)
		}
		counter++
	}
	z01.PrintRune('\n')
}
