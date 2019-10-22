package student

func IsALower(str string) bool {
	runes := []rune(str)
	ok := false
	for _, i := range runes {
		if !(i >= 'a' && r <= 'z') {
			return ok
		}
		ok = true
	}
	return ok
}
