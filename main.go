package main

import "fmt"

func SwapBits(octet byte) byte {
	return octet<<4 | octet>>4
}

func main() {
	var num byte = 65
	fmt.Printf("%08b\n", num)
	fmt.Printf("%08b\n", SwapBits(num))
}
