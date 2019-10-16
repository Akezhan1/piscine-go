package printcomb

import "github.com/01-edu/z01"

func PrintComb() {
	var a,b,c rune = '0','1','2'
	for i := a; i <= 54; i++ {
		for k := b; k <= 56; k++ {
			for j := c; j <= 57; j++ {
				if i < k {
					if k < j {
						z01.PrintRune(i)
						z01.PrintRune(k)
						z01.PrintRune(j)
						z01.PrintRune(44)
						z01.PrintRune(32)
					}
				}
			}
		}
	}
	z01.PrintRune(55)
	z01.PrintRune(56)
	z01.PrintRune(57)
}