package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	params := os.Args
	for _, value := range params[1:] {
		for _, v := range value {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
