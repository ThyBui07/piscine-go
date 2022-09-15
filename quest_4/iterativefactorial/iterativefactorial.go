package piscine

func IterativeFactorial(nb int) int {
	res := 1
	// if over 20 => not fit in integer
	if nb < 0 || nb > 20 {
		res = 0
	} else if nb == 0 {
		res = 1
	} else {
		for i := 1; i <= nb; i++ {
			res = res * i
		}
	}
	return res
}
