package student

func TrimAtoi(s string) int {
	if s == "" {
		return 0
	}
	x := 0
	k := 1
	s1 := AlphaCount(s)

	if s != "" {
		for i, n := range s1 {
			y := 0
			if n < '0' || n > '9' {
				if n == '-' || n == '+' {
					if i != 0 {
						return 0
					}
					if n == '-' {
						k = -1
					}
				} else {
					return 0
				}
			}
			for i := '1'; i <= n; i++ {
				y++
			}
			x = x*10 + y
		}
	}
	return x * k
}

func AlphaCount(str string) []rune {
	//result := 0
	len := 0
	var res []rune
	var final []rune
	for i := range str {
		len = i
	}
	for i := 0; i <= len; i++ {
		if str[i] >= '0' && str[i] <= '9' || str[i] == '-' {
			res = append(res, rune(str[i]))
		}
	}
	for i := range res {
		if res[i] == '-' {
			if i == 0 {
				return res
			} else {
				for i := 0; i <= len; i++ {
					if str[i] >= '0' && str[i] <= '9' {
						final = append(final, rune(str[i]))
					}
				}
			}
			return final
		}
	}
	return res
}
