package student

func Capitalize(s string) string {
	srunes := []rune(s)
	len := 0
	for i := range srunes {
		len = i
	}
	ok := false
	for i := 0; i < len; i++ {
		if trust(srunes[i]) == true && ok {
			if srunes[i] >= 'a' && srunes[i] <= 'z' {
				srunes[i] = ToUpper(srunes[i])
			}
			ok = false
		} else if srunes[i] >= 'A' && srunes[i] <= 'Z' {
			if i == 0 {
				continue
			}
			srunes[i] = ToLower(srunes[i])
		} else if trust(srunes[i]) == false {
			ok = true
		}
	}
	return string(srunes)
}

func trust(a rune) bool {
	if (a >= 'A' && a <= 'Z') || (a >= 'a' && a <= 'z') || (a >= '0' && a <= '9') {
		return true
	}
	return false
}

func ToUpper(s rune) rune {
	return rune(s - 32)
}

func ToLower(s rune) rune {
	return rune(s + 32)
}
