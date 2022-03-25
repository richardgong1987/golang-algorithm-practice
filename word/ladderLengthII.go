package word

//func ladderLengthII(beginWord string, endWord string, wordArray []string) [][]string {
//	reached := NewStringSet([]string{beginWord})
//	wordList := NewStringSet(wordArray)
//	distance := 1
//
//	for !reached.Contains(endWord) {
//		toAdd := NewStringSet([]string{})
//
//		for each := range reached.M {
//			for i, _ := range each {
//				chars := []byte(each)
//				for ch := 'a'; ch <= 'z'; ch += 1 {
//					chars[i] = byte(ch)
//					word := string(chars)
//					if wordList.Contains(word) {
//						toAdd.Add(word)
//						wordList.Remove(word)
//					}
//				}
//			}
//		}
//
//		distance += 1
//		if toAdd.Len() == 0 {
//			return 0
//		}
//		reached = toAdd
//	}
//
//	return distance
//}
