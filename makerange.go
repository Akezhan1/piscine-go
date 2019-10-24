package student

func MakeRange(min, max int) []int {
	if min > max || min == max {
		return []int(nil)
	}
	len := max - min
	res := make([]int, len)
	for k := 0; k < len; k++ {
		for i := min; i < max; i++ {
			res[k] = min + k
		}
	}
	return res
}
