package word

import "container/list"

func ladderLengthII(beginWord string, endWord string, wordList []string) [][]string {
	var dict = make(map[string]struct{}, 0)
	for _, dictWord := range wordList {
		dict[dictWord] = struct{}{}
	}
	depth := markDepth(beginWord, endWord, dict)
	result := make([][]string, 0)
	wordLadderHelper(depth, &result, beginWord, beginWord, endWord, dict, []string{})
	return result
}

func markDepth(beginWord string, endWord string, dict map[string]struct{}) map[string]int {
	queue := list.New()
	depth := make(map[string]int)

	queue.PushBack(beginWord)
	depth[beginWord] = 1

	for queue.Len() > 0 {
		ele := queue.Front().Value.(string)

		queue.Remove(queue.Front())

		for _, word := range nextWords(dict, ele) {
			if _, exist := depth[word]; !exist {
				depth[word] = depth[ele] + 1
				queue.PushBack(word)
				if word == endWord {
					return depth
				}
			}
		}
	}
	return nil
}

func wordLadderHelper(depth map[string]int, results *[][]string, current string, startWord string, endWord string, dict map[string]struct{}, result []string) {
	if current == endWord {
		tmp := make([]string, len(result), len(result)+1)
		copy(tmp, result)
		tmp = append(tmp, endWord)
		*results = append(*results, tmp)
	}

	for _, word := range nextWords(dict, current) {
		if depth[word] != depth[current]+1 {
			continue
		}
		result = append(result, current)
		wordLadderHelper(depth, results, word, startWord, endWord, dict, result)
		result = result[:len(result)-1]
	}
}

func nextWords(dict map[string]struct{}, word string) []string {
	nextWordList := make([]string, 0)
	for i := 0; i < len(word); i++ {
		// loop 26 letters
		for j := 0; j < 26; j++ {
			if word[i] == byte(int('a')+j) {
				continue
			}
			tmp := word[:i] + string(rune(int('a')+j)) + word[i+1:]
			if containsStr(dict, tmp) {
				nextWordList = append(nextWordList, tmp)
			}
		}
	}
	return nextWordList
}

func containsStr(dict map[string]struct{}, target string) bool {
	_, exist := dict[target]
	return exist
}
