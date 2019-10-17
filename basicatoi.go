package basicatoi
func BasicAtoi(s string) {
	runes := []rune(s)
	var final int
	var res []int
	n := '0'
	v := 0
		res = append(res, v)
	for i, c := range runes {
		for n != c {
			n++
			v++
		}
		res = append(res, v)
	}
	for i := range runes {
		final = final*10 + res[i]
	}
}
