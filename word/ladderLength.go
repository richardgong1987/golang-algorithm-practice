package word

/**
Input: beginWord = "hit", endWord = "cog", wordList = ["hot","dot","dog","lot","log","cog"]
*/
func ladderLength(beginWord string, endWord string, wordArray []string) int {
	reached := NewStringSet([]string{beginWord})
	wordList := NewStringSet(wordArray)
	distance := 1

	for !reached.Contains(endWord) {
		toAdd := NewStringSet([]string{})

		for each := range reached.M {
			for i, _ := range each {
				chars := []byte(each)
				for ch := 'a'; ch <= 'z'; ch += 1 {
					chars[i] = byte(ch)
					word := string(chars)
					if wordList.Contains(word) {
						toAdd.Add(word)
						wordList.Remove(word)
					}
				}
			}
		}

		distance += 1
		if toAdd.Len() == 0 {
			return 0
		}
		reached = toAdd
	}

	return distance
}

type StringSet struct {
	M map[string]bool
}

func NewStringSet(str []string) *StringSet {
	set := &StringSet{M: map[string]bool{}}
	set.AddAll(str)
	return set
}
func (set *StringSet) Add(e string) bool {
	if !set.M[e] {
		set.M[e] = true
		return true
	}
	return false
}
func (set *StringSet) AddAll(e []string) {
	for _, value := range e {
		set.M[value] = true
	}
}
func (set *StringSet) Remove(e string) {
	delete(set.M, e)
}

func (set *StringSet) Contains(e string) bool {
	return set.M[e]
}

//获取元素数量
func (set *StringSet) Len() int {
	return len(set.M)
}
