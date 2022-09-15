package piscine

func anyAlpha(n rune) bool {
	if (n >= 'A' && n <= 'Z') || (n >= 'a' && n <= 'z') || (n >= '0' && n <= '9') {
		return true
	}
	return false
}

func Capitalize(s string) string {
	blank := true
	c := []rune(s)
	for i := 0; i < len(s); i++ {
		if blank && anyAlpha(c[i]) {
			if c[i] >= 'a' && c[i] <= 'z' {
				c[i] = c[i] + 'A' - 'a'
			}
			blank = false
		} else if c[i] >= 'A' && c[i] <= 'Z' {
			c[i] = c[i] + 'a' - 'A'
		} else if !anyAlpha(c[i]) {
			blank = true
		}
	}
	return string(c)
}
