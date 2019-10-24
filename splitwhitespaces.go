package student

func SplitWhiteSpaces(str string) []string {
	temp := ""
	res := []string{}
	len := 0
	for i := range str {
		len = i
	}
	for i := 0; i < len+1; i++ {
		if i == len && string(str[i]) != " " && string(str[i]) != "\t" && string(str[i]) != "\n" {
			temp += string(str[i])
			res = append(res, temp)
		} else if string(str[i]) != " " && string(str[i]) != "\t" && string(str[i]) != "\n" {
			temp += string(str[i])
		} else {
			res = append(res, temp)
			temp = ""
		}
	}
	return res
}
