package main

import "github.com/01-edu/z01"

func main() {
	var a rune = '\n'
	for i := 'z'; i <= 'a'; i-- {
		z01.PrintRune(i)
	}
	z01.PrintRune(a)
}
