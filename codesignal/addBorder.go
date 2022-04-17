package codesignal

import "strings"

func AddBorder(picture []string) []string {
	startEnd := strings.Repeat("*", len(picture[0])+2)
	list := []string{startEnd}
	for _, s := range picture {
		list = append(list, "*"+s+"*")
	}
	return append(list, startEnd)
}
