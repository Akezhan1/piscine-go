package main

import "github.com/01-edu/z01"

func PrintComb() {
	var a, b, c, d rune = '0', '0', '0', '1'
	for i := c; i <= 57; i++ {
		for k := d; k <= 57; k++ {
				if i <= k {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(32)
						z01.PrintRune(i)
						z01.PrintRune(k)
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
}