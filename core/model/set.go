package model

// Set use struct{} to represent a set
type Set map[int64]struct{}

func NewSet(slice []int64) Set {
	set := make(Set)
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}
func (set Set) Contains(v int64) bool {
	_, ok := set[v]
	return ok
}
func (set Set) Add(v int64) {
	set[v] = struct{}{}
}
func (set Set) Remove(v int64) {
	delete(set, v)
}
func (set Set) Clear() {
	set.Clear()
}
func (set Set) Size() int64 {
	return int64(len(set))
}
