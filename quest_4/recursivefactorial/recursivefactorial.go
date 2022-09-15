package piscine

func RecursiveFactorial(nb int) int {
	result := 1
	if nb < 0 || nb > 20 {
		result = 0
	} else if nb == 0 {
		result = 1
	} else {
		result = nb * RecursiveFactorial(nb-1)
	}
	return result
}
