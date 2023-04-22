package treex

type defaultTree struct {
	key       string
	val       any
	parentKey string
	parent    TreeNode
	children  []TreeNode
}

func IsTopNode(n TreeNode) bool {
	return n.GetKey() == ""
}

func NewDefaultTree(key, parent string, val any) TreeNode {
	return &defaultTree{key: key, val: val, parentKey: parent}
}

func (t *defaultTree) GetKey() string {
	return t.key
}
func (t *defaultTree) GetValue() any {
	return t.val
}
func (t *defaultTree) GetParent() TreeNode {
	return t.parent
}
func (t *defaultTree) GetChildren() []TreeNode {
	return t.children
}
func (t *defaultTree) AppendChild(n TreeNode) {
	if t.children == nil {
		t.children = make([]TreeNode, 0)
	}
	t.children = append(t.children, n)
}
func (t *defaultTree) BindingParent(n TreeNode) {
	t.parent = n
}

func (t *defaultTree) IsTop() bool {
	return t.parentKey == ""
}
func (t *defaultTree) GetParentKey() string {
	return t.parentKey
}
