package util

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
