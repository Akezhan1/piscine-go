package strlen

func StrLen(str string) int {
	strrune := []rune(str)
	var a int
	for i := range strrune; i++ {
		a = i
	}
	return a
}
