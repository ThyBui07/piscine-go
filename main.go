package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) != 0 {
		str := args[len(args)-1]
		for _, i := range str {
			z01.PrintRune(i)

		}
	} else {
		os.Exit(0)
	}
	z01.PrintRune('\n')
}
