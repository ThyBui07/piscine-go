package piscine

func RecursivePower(nb int, power int) int {
	if nb == 0 && power == 0 {
		return 1
	} else if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else if power == 1 {
		return nb
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}
