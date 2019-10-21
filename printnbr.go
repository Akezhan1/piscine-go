package student
import "github.com/01-edu/z01"

func ToString(n int) string {
	buf := [11]byte{}
	pos := len(buf)
	i := n
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}

func PrintNbr(n int) {
	a := ToString(n)
	for _, val := range a {
		z01.PrintRune(val)
	}
}
