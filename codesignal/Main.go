package codesignal

import (
	"fmt"
	"math"
)

func Main() {
	i := solution([]int{2, 4, 7})
	fmt.Println(i)
}

func solution(a []int) int {
	arrLen := len(a)
	if arrLen <= 1 {
		return 0
	}

	min := math.MaxFloat64
	var ret int
	for i := 0; i < arrLen; i++ {
		var innerSum float64
		for j := 0; j < arrLen; j++ {
			innerSum += math.Abs(float64(a[i] - a[j]))
		}

		if innerSum < min {
			min = innerSum
			ret = a[i]
		}
	}
	return ret

}
