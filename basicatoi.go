package main

func BasicAtoi(s string) {
	runes := []rune(s)
	var final int
	var res []int = make([]int, len(runes))
	n := '0'
	v := 0
	for i ,c:= range runes {
		for n != c {
		n++
		v++
		}
		res[i] = v
	}

	for i := range runes {
		final = final*10 + res[i]
	}

	fmt.Println(final)
}