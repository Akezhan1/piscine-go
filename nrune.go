package student

func NRune(s string, n int) rune {
	runes := []rune(s)
	len := 0
	final := 0
	for i := range runes {
		len = i
	}
	if n > len {
		return 0
	}
	for i := 1; i < n; i++ {
		final = n
	}
	return runes[final]
}
