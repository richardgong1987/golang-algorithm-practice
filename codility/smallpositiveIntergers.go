package codility

import "sort"

func SmallpositiveIntergers(A []int) int {
	// write your code in Go 1.4
	sort.Ints(A)
	maxv := A[len(A)-1]

	if maxv <= 0 {
		return 1
	}

	for i := 1; i < maxv+1; i++ {
		if !check(A, i) {
			return i
		}
	}

	return maxv + 1
}
func check(A []int, b int) bool {

	for _, n := range A {
		if n == b {
			return true
		}
	}

	return false
}
