package codesignal

import (
	"fmt"
	"strconv"
)

func Main() {
	fmt.Println(FileNaming([]string{
		"a(1)",
		"a(6)",
		"a",
		"a"}))
}

func solution(names []string) []string {
	m := map[string]int{}
	var str []string
	for _, name := range names {
		result, exist := m[name]
		if exist {
			if result == 1 {
				key := name + "(" + strconv.Itoa(result) + ")"
				result2, exist2 := m[key]
				if exist2 {
					key2 := name + "(" + strconv.Itoa(result2+1) + ")"
					m[key2]++
					m[name]++
					str = append(str, key2)
				} else {
					m[key]++
					m[name]++
					str = append(str, key)

				}
			} else if result > 1 {
				key := name + "(" + strconv.Itoa(result) + ")"
				m[key]++
				m[name]++
				str = append(str, key)
			}
		} else {
			str = append(str, name)
			m[name]++
		}
	}
	return str
}
