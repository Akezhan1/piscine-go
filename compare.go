package student

func compare(a, b string) int {
	arunes := []rune(a)
	brunes := []rune(b)
	count := 0
	len := 0
	for i := range a {
		len = i + 1
	}
	for k := range brunes {
		if IndexRune(arunes, brunes[k]) {
			count++
		}
	}
	switch {
	case len == count:
		return 0
	case len < count:
		return -1
	case len > count:
		return 1
	}
	return len
}

func IndexRune(s []rune, r rune) bool {
	indexrune := false
	for i := range s {
		if s[i] == r {
			indexrune = true
			return indexrune
		}
	}
	return indexrune
}
