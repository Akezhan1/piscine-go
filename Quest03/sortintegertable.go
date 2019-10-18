package main

import "fmt"

func SortIntegerTable(table []int) {
	sort := false
	var b int
	for i := range table {
		b = i
	}
	b = b+1
	for !sort {
		swap := false
		for i := 0; i < b-1; i++ {
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

func main() {
	i := []int{-1,5,4,-8,5,-3}
	SortIntegerTable(i)
	fmt.Println(i)
}