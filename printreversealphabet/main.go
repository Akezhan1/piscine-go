package main

import "github.com/01-edu/z01"

func main() {
	var a rune = 'z'
	for i := a; i <= 'a'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}
