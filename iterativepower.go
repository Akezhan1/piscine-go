package student

func IterativePower(nb int, power int) int {
	res := 1
	for i := 1; i <= power; i++ {
		res = nb * res
	}
	if res < 0 {
		return 0
	}
	return res
}
