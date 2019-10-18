package atoi

func Atoi(s string) int {
	res := 0
	for i, val := range s {
		if s[0] >= '0' && s[0] <= '9' || s[0] == '+' {
			if val == 32 {
				return 0
			}
			a := 0
			for i := '1'; i <= val; i++ {
				a++
			}
			res = res*10 + a
		} else if s[0] == '-' {
			a := 0
			for i := '1'; i <= val; i++ {
				a--
			}
			res = res*10 + a
		}
		if i > 0 {
			if !(val >= '0' && val <= '9') {
				return 0
			}
		}
	}
	return res
}
