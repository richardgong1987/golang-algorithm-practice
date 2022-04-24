package codesignal

func StringsRearrangement(inputArray []string) bool {
	inputLen := len(inputArray)

	for i, iStr := range inputArray {
		tmp := []string{iStr}
		for j, jStr := range inputArray {
			lastOne := tmp[len(tmp)-1]
			if i != j && exactOneDiff(lastOne, jStr) {
				tmp = append(tmp, jStr)
			}
		}
		if inputLen-1 == len(tmp) {
			return true
		}
	}
	return false
}
func exactOneDiff(a string, b string) bool {
	aLen := len(a)
	diff := 0
	for i := 0; i < aLen; i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff == 1
}
