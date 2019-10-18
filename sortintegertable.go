package sortintgertable

func SortIntegerTable(table []int) {
	sort := false
	var b int
	for i := range table {
		b = i
	}
	for !sort {
		swap := false
		for i := 0; i < b+1; i++ {
			if table[i+1] < table[i] {
				table[i+1], table[i] = table[i], table[i+1]
				swap = true
			}
		}
		if !swap {
			sort = true
		}
	}
}
