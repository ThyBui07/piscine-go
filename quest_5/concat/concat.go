package piscine

func Concat(str1 string, str2 string) string {
	array1 := []rune(str1)
	array2 := []rune(str2)
	arrayCom := []rune{}

	arrayCom = append(arrayCom, array1...)
	arrayCom = append(arrayCom, array2...)

	return string(arrayCom)
}
