package student

func isprime(nb int) bool {
	if nb == 1 {
		return false
	}
	if nb <= 0 {
		return false
	} else {
		for i := 2; i*i <= nb; i++ {
			if nb%i == 0 {
				return false
			}
		}
		return true
	}
}

func FindNextPrime(n int) int {
	for i := 0; i <= n; i++ {
		if isprime(n) == false {
			n++
		} else {
			break
		}
	}
	return n
}
