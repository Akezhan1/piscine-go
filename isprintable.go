package student

func IsPrintable(str string) bool {
	runes := []rune(str)
	ok := true
	for _, i := range runes {
		if !(i >= 'A' && i <= 'Z') || !(i >= 'a' && i <= 'z') {
			return ok
		}
		ok = false
	}
	return ok
}
