package piscine

func Swap(a *int, b *int) {
	c := *a
	d := *b
	*a = d
	*b = c
}
