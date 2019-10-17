package main
func Swap(a *int, b *int) {
	d := *a
	c := *b
	*a = c
	*b = d
}
