package student

import (
	"github.com/01-edu/z01"
)

func ToMassive(n int) []int {
	var res []int
	for n > 0 {
		if n == 0 {
			res = append(res, n)
		} else {
			a := n % 10
			res = append(res, a)
		}
		n = n / 10
	}
	return res
}

func PrintNbrInOrder(n int) {
	if n < 0 {
	} else if n == 0 {
		z01.PrintRune('0')
	} else {
		mod := '0'
		for _, i := range ToMassive(n) {
			for k := 0; k < i; k++ {
				mod++
			}
			z01.PrintRune(mod)
			mod = '0'
		}
	}
}
