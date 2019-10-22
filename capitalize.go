package student

func Capitalize(s string) string {
	srunes := []rune(s)
	for i := 0; i < len(s); i++ {
		if srunes[i] == ' ' || srunes[i] == '+' || srunes[i] == '-' {
			temp := srunes[i+1]
			srunes[i+1] = ToUpper(rune(temp))
		} else {
			srunes[i] = srunes[i]
		}
	}
	return string(srunes)
}

func ToUpper(s rune) rune {
	if rune(s-32) < 65 {
		return s
	}
	return rune(s - 32)
}
