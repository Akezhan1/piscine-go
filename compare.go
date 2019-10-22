package main

<<<<<<< HEAD
import "fmt"

=======
>>>>>>> 5f92a4b19cc66ebbcf532ec0932596489a9e7172
func Compare(a, b string) int {
	arunes := []rune(a)
	brunes := []rune(b)
	count := 0
	len := 0
	for i := range a {
		len = i + 1
	}
	for i := range brunes {
		if IndexRune(arunes, brunes[i]) {
			count++
		}
	}
	if len == count {
		return 0
<<<<<<< HEAD
	}

	for i := range a {
		for k := range b {
			switch {
			case a[i] > b[k]:
				return 1
			case a[i] < b[k]:
				return -1
			}
		}
=======
	case len < count:
		return 1
	case len > count:
		return 1
>>>>>>> 5f92a4b19cc66ebbcf532ec0932596489a9e7172
	}
	return len
}

func IndexRune(s []rune, r rune) bool {
	indexrune := false
	for i := range s {
		if s[i] == r {
			indexrune = true
			return indexrune
		}
	}
	return indexrune
}

func main() {
	fmt.Println(Compare("Salut!", "lut!"))
}
