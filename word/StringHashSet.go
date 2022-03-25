package word

type StringHashSet map[string]bool

func NewStringHashSet(es []string) StringHashSet {
	set := StringHashSet{}
	set.AddALL(es)
	return set
}
func (s StringHashSet) Len() int {
	return len(s)
}
func (s StringHashSet) Add(e string) bool {
	if !s[e] {
		s[e] = true
		return true
	}
	return false
}
func (s StringHashSet) AddALL(e []string) {
	for _, element := range e {
		s[element] = true
	}
}
func (s StringHashSet) Clear() {
	for key := range s {
		delete(s, key)
	}
}
func (s StringHashSet) Remove(e string) {
	delete(s, e)
}
func (s StringHashSet) Contains(e string) bool {
	return s[e]
}
func (s StringHashSet) isEmpty() bool {
	return len(s) == 0
}
