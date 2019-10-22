package student

func IsLower(str string) bool {
	runes := []rune(str)
	ok := true
	for _, i := range runes {
		if !(i >= 'a' && i <= 'z') {
			ok = false
			return ok
		}
		ok = true
	}
	return ok
}
