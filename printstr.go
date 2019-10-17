package printstr
import "github.com/01-edu/z01"

func PrintStr(str string) {
	strrune := []rune(str)
	for i := 0; i < len(strrune); i++ {
		z01.PrintRune(strrune[i])
	}
	z01.PrintRune(10)
}
