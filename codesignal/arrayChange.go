package codesignal

func ArrayChange(inputArray []int) int {

	max := inputArray[0]
	moves := 0

	for i := 1; i < len(inputArray); i++ {
		if inputArray[i] <= max {
			moves += max - inputArray[i] + 1
			inputArray[i] = max + 1
		}
		max = inputArray[i]
	}
	return moves

}
