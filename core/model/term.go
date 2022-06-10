package model

type Term struct {
	Index *Index
	Score float64
}

// Terms is a small heap
type Terms []*Term

func (Terms *Terms) Size() int64 {
	return int64(len(*Terms))
}
func (Terms *Terms) Less(i, j int) bool {
	return (*Terms)[i].Score < (*Terms)[j].Score
}
func (Terms *Terms) Swap(i, j int) {
	(*Terms)[i], (*Terms)[j] = (*Terms)[j], (*Terms)[i]
}
func (Terms *Terms) Push(Term *Term) {
	*Terms = append(*Terms, Term)
}
func (Terms *Terms) Pop() (*Term, bool) {
	if len(*Terms) == 0 {
		return nil, false
	}
	Term := (*Terms)[len(*Terms)-1]
	*Terms = (*Terms)[:len(*Terms)-1]
	return Term, true
}
func (Terms *Terms) Empty() bool {
	return Terms.Size() == 0
}
