package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	lenIndex := 0
	letterBool := false
	letterIndex := 96
	for index, key := range os.Args {
		lenIndex = index
		if key == "---upper" {
			letterBool = true
		}
	}
	if letterBool == true {
		letterIndex = 64
	}
	for i := 1; i <= lenIndex; i++ {
		for _, key := range os.Args[i] {
			z01.PrintRune(key + rune(letterIndex))
		}
		z01.PrintRune(10)
	}
}
