package student

func IterativeFactorial(y int) int {
	res := 1
	if res < 0 || res > 50 {
		return 0
	} else if y == 0 {
		return 0
	} else {
		for i := 1; i <= y; i++ {
			res = i * res
		}
	}
	return res
}
