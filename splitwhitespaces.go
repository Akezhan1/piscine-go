package student

func SplitWhiteSpaces(str string) []string {
	len := 0
	len2 := 0
	for _, i := range str {
		if i == ' ' || i == '\t' || i == '\n' {
			len++
		}
	}
	for i := range str {
		len2 = i
	}
	res := make([]string, len+1)
	temp := ""
	check := true
	k := 0
	for i := 0; i <= len2; i++ {
		if i == len2 {
			temp += string(str[i])
			res[k] = temp
		} else if str[i] == ' ' || str[i] == '\t' || str[i] == '\n' {
			res[k] = temp
			temp = ""
			check = false
			if !check {
				k++
			}
			check = true
			continue
		} else {
			temp += string(str[i])
		}
	}
	return res
}
