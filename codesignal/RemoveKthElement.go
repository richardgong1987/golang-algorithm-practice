package codesignal

func RemoveKthElement(inputArray []int, k int) []int {
	res := []int{}

	for i := 0; i < len(inputArray); i++ {
		if (i+1)%k != 0 {
			res = append(res, inputArray[i])
		}
	}

	return res
}
