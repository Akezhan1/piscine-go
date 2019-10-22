package student

func IsUpper(str string) bool {
	runes := []rune(str)
	ok := true
	for _, i := range runes {
		if !(i >= 'A' && i <= 'Z') {
			ok = false
			return ok
		}
	}
	return ok
}
