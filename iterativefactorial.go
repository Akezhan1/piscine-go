package student

func IterativeFactorial(y int) int {
	res := 1
	if y < 0 || y > 50 {
		return 0
	} else if y == 0 {
		return 1
	} else {
		for i := 1; i <= y; i++ {
			res = i * res
		}
	}
	if res < 0 {
		return 0
	}
	return res
}
