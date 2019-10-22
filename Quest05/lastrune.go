package student

func LastRune(s string) rune {
	runes := []rune(s)
	last := 0
	for i := range runes {
		last = i
	}
	return runes[last]
}
