package codesignal

import "strconv"

func FileNaming(names []string) []string {
	for i, name := range names {
		if Contains(names[:i], name) {
			j := 1
			for Contains(names[:i], names[i]+"("+strconv.Itoa(j)+")") {
				j++
			}
			names[i] += "(" + strconv.Itoa(j) + ")"
		}
	}
	return names
}

func Contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
