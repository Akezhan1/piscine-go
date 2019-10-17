package strrev
func StrRev(str string) string {
	revstr := []rune(str)
	strrune := []rune(str)
	var a int

	for i := range strrune {
		a = i
	}

	for i := range revstr {
		revstr[i] = strrune[a]
		a--
	}
	return string(revstr)
}
