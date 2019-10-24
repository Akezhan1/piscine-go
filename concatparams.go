package student

func ConcatParams(args []string) string {
	var res string
	for i, val := range args {
		if i == 0 {
			res = val
			continue
		}
		res += '\n' + val
	}
	return res
}
