package main

import "github.com/01-edu/z01"

func main() {
	j := 'A'
	j++
	for i:= 'a'; i <= 'z'; i= i+2 {
		z01.PrintRune(i)
		for k:=j; k <= 'Z'; k++ {
			z01.PrintRune(j)
			j = j + 2
			break
		}
	}
	z01.PrintRune('\n')
}