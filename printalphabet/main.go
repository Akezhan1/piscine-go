package main 

import (
	"z01"
)

func main(){
	for i:= 'a'; i<='z'; i++{
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
	
}
