package student

func ToLower(s string) string {
	srunes := []rune(s)
	for i := 0; i < len(s); i++ {
		if srunes[i] >= 'A' && srunes[i] <= 'Z' {
			srunes[i] = rune(srunes[i] + 32)
		}
	}
	return string(srunes)
}
