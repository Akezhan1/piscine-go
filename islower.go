package student

func IsLower(str string) bool {
	len := 0
	count := 0
	for i := range str {
		len = i
	}
	for i := 0; i <= len; i++ {
		if str[i] >= 0 && str[i] <= 64 || str[i] >= 91 && str[i] <= 96 || str[i] >= 123 && str[i] <= 127 {
			return false
		} else {
			if lowercase(rune(str[i])) {
				count++
			} else {
				count = 0
			}
		}
	}
	if count-1 == len {
		return true
	}
	return false
}

func lowercase(s rune) bool {
	if s+32 > 122 {
		return true
	}
	return false
}
