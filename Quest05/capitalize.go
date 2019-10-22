package student

func trust(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func ToUpper(s rune) rune {
	if s-32 < 65 {
		return s
	}
	return rune(s - 32)
}

func ToLower(s rune) rune {
	if s+32 > 122 {
		return s
	}
	return rune(s + 32)
}

func Capitalize(s string) string {
	srunes := []rune(s)
	len := 0
	ok := true
	for i := range s {
		len = i
	}

	for i := 0; i < len; i++ {
		if trust(srunes[i]) == true && ok {
			if srunes[i] >= 'a' && srunes[i] <= 'z' {
				srunes[i] = ToUpper(srunes[i])
			} else if srunes[i] >= '0' && srunes[i] <= '9' {
				srunes[i+1] = ToLower(srunes[i+1])
			}
			ok = false
		} else if srunes[i] >= 'A' && srunes[i] <= 'Z' {
			srunes[i] = ToLower(srunes[i])
		} else if trust(srunes[i]) == false {
			ok = true
		}
	}
	return string(srunes)
}
