package tree

type BasicNode interface {
	AppendChild(BasicNode)
}

type ExtendNode interface {
	Sort
	GetKey() string
	AppendChild(ExtendNode)
}

type exNode struct {
	key      string
	weight   int
	children []ExtendNode
	keyMap   map[string]bool
}
