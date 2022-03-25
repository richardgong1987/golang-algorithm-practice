package word

import (
	"fmt"
)

func Main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	length := ladderLength(beginWord, endWord, wordList)
	fmt.Println(length)
}
