package word

func ladderLength(beginWord string, endWord string, wordArray []string) int {
	wordList := NewStringHashSet(wordArray)
	searchList := NewStringHashSet([]string{beginWord})
	sum := 1

	for !searchList.Contains(endWord) {
		bucket := NewStringHashSet([]string{})
		for hitWord := range searchList {
			for i := range hitWord {
				chars := []byte(hitWord)
				for ch := 'a'; ch <= 'z'; ch++ {
					chars[i] = byte(ch)
					searchWord := string(chars)
					if wordList.Contains(searchWord) {
						bucket.Add(searchWord)
						wordList.Remove(searchWord)
					}
				}
			}
		}

		sum++
		if bucket.isEmpty() {
			return 0
		}
		searchList = bucket
	}
	return sum
}
