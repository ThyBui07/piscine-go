package main

import "fmt"

func main() {
	a := []int{23, 123, 1, 11, 55, 93}
	max := Max(a)
	fmt.Println(max)
}

func Max(a []int) int {
	max := a[0]
	for i := 0; i < len(a)-1; i++ {
		if a[i] < a[i+1] && a[i+1] > max {
			max = a[i+1]
		}
	}
	return max
}
