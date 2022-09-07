package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	str := ""
	if len(args) < 2 {
		z01.PrintRune('\n')
	} else {
		str = args[0] + args[1]
		str = RemoveDup(str)
		for _, x := range str {
			z01.PrintRune(x)
		}
	}
	z01.PrintRune('\n')
}

func RemoveDup(s string) string {
	res := ""

	for _, x := range s {
		var dup bool
		for _, y := range res {
			if x != y {
				dup = false
			} else {
				dup = true
				break
			}
		}
		if dup == false {
			res += string(x)
		}
	}
	return res
}
