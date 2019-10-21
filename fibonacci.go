package student

func Fibonacci(x int) int {
	res := 0
	if x < 0 {
		return -1
	} else if x < 2 {
		return x
	} else {
		res = (Fibonacci(x-1) + Fibonacci(x-2))
	}
	return res
}
