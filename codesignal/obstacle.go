package codesignal

func AvoidObstacles(inputArray []int) int {
	set := map[int]bool{}

	var maxV int
	for _, v := range inputArray {
		set[v] = true
		if v > maxV {
			maxV = v
		}
	}

	for i := 1; i <= maxV; i++ {
		j := i
		for ; j <= maxV; j += i {
			if set[j] {
				break
			}
		}
		if j > maxV {
			return i
		}
	}

	return maxV + 1
}
