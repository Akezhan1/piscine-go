package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	len := 0
	sort := false
	for i := range arguments {
		len = i
	}
	for i := 1; i <= len; i++ {
		for _, k := range arguments[i] {
			for !sort {
				swap := false
				for i := 0; i < len; i++ {
					if arguments[i+1] < arguments[i] {
						arguments[i+1], arguments[i] = arguments[i], arguments[i+1]
						swap = true
					}
				}
				if !swap {
					sort = true
				}
			}
			z01.PrintRune(k)
		}
		z01.PrintRune(10)
	}
}

