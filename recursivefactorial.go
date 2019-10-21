package piscine

func RecursiveFactorial(nb int) int {
	res := 1
	if nb < 0 || nb > 50 {
		return 0
	} else if nb == 0 {
		return 1 
	} else if nb == 1 {
		return 1
	} else nb > 1 { 
		res = nb * RecursiveFactorial(nb - 1)
	}

	if res < 0 {
		return 0
	}

	return res
}