package student

func ToUpper(s string) string {
	var res byte
	var srunes []rune
	var final string
	//var byt byte = 32
	for i := range s {
		res = s[i] - 32
		if res < 65 {
			res += 32
			srunes = append(srunes, rune(res))
		} else {
			srunes = append(srunes, rune(res))
		}
	}
	final = string(srunes)
	return final
}
