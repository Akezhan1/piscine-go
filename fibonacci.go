package student

func FIbonacci(x int) int {
	res := 0
	if x < 0 {
		return -1
	} else if x < 2 {
		return x
	} else {
		res = (fib(x-1) + fib(x-2))
	}
	return res
}
