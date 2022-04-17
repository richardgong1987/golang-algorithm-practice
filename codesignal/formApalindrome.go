package codesignal

func FormPalindrome(inputString string) bool {
	letter26 := [256]rune{}

	for _, i := range inputString {
		letter26[i]++
	}

	odd := 0

	for _, c := range letter26 {
		if c&1 == 1 {
			odd++
		}
		if odd > 1 {
			return false
		}
	}

	return false

}
