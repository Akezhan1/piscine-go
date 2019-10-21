package piscine

func IterativeFactorial(y int) int {
	res := 1
	if y == 0 {
		return 0
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
