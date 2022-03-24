package util

func sliceContains(slice []string, word string) (index int, isExit bool) {
	for i, s := range slice {
		if s == word {
			return i, true
		}
	}
	return -1, false
}
func sliceRemove(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}
