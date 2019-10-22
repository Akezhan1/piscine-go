package student

func NRune(s string, n int) rune {
	runes := []rune(s)
	len := 0
	final := 0
	for i := range runes {
		len = i + 1
	}
	if n > len {
		return 0
	} else if n < 0 {
		return runes[0]	
	}
	for i := 0; i < n; i++ {
		final = i
	}
	return runes[final]
}
