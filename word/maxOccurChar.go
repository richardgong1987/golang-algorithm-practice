package word

import (
	"sort"
	"strconv"
)

func MaxLenChar(a []int) []int {
	m := map[int]int{}

	for i := 0; i < 10; i++ {
		m[i] = 0
	}

	for _, value := range a {
		str := strconv.FormatInt(int64(value), 10)
		for _, c := range str {
			v, _ := strconv.Atoi(string(c))
			m[v]++
		}
	}
	maxValue := -1

	for _, v := range m {
		if v > maxValue {
			maxValue = v
		}
	}

	ret := []int{}

	for i, v := range m {
		if m[v] == maxValue {
			ret = append(ret, i)
		}
	}
	sort.Ints(ret)
	return ret
}
