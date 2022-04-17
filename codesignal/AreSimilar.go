package codesignal

import (
	"reflect"
	"sort"
)

func AreSimilar(a []int, b []int) bool {
	diffCnt := 0

	for i := range a {
		if a[i] != b[i] {
			diffCnt++
		}
	}
	sort.Ints(a)
	sort.Ints(b)
	if reflect.DeepEqual(a, b) {
		return diffCnt <= 2
	}
	return false
}
