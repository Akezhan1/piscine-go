package student

func IsAlpha(str string) bool {
	for i := range str {
		if str[i] >= 0 && str[i] <= 47 ||
			str[i] >= 58 && str[i] <= 64 ||
			str[i] >= 91 && str[i] <= 96 ||
			str[i] >= 123 && str[i] <= 127 {
			return false
		}
	}
	return true
}
