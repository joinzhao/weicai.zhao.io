package tree

type Sort interface {
	Weight() int
}

type SortNode interface {
	BasicNode
	Sort
}

type SortNodes []SortNode

func (s SortNodes) Len() int           { return len(s) }
func (s SortNodes) Less(i, j int) bool { return s[i].Weight() > s[j].Weight() }
func (s SortNodes) Swap(i, j int)      { s[i] = s[j] }
