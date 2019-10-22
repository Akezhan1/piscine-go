package student

func IsNumeric(str string) bool {
	for i := range str {
		if str[i] >= 0 && str[i] <= 47 ||
			str[i] >= 58 && str[i] <= 127 {
			return false
		}
	}
	return true
}
