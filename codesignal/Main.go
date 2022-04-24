package codesignal

import (
	"fmt"
)

func Main() {
	fmt.Println(solution("010010000110010101101100011011000110111100100001"))
}

func solution(code string) string {
	codeLen := len(code)
	str := ""
	for i := 0; i <= codeLen-8; i += 8 {
		str += string(BinarrytoInt(code[i : i+8]))
	}
	return str
}
