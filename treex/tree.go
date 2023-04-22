package treex

// Node 普通节点， 键值对
type Node interface {
	GetKey() string
	GetValue() any
}

// TreeNode 树节点， 拥有普通节点的同时，拥有父节点和子节点
type TreeNode interface {
	Node
	GetParent() TreeNode
	GetChildren() []TreeNode
	AppendChild(TreeNode)
	BindingParent(TreeNode)
	GetParentKey() string
	IsTop() bool
}

// SimpleTree 多课树
type SimpleTree struct {
	// 根节点集合
	tops []TreeNode
}

// Foreach 树节点遍历
func (t *SimpleTree) Foreach(f func(TreeNode)) {
	if t.tops != nil {
		for _, top := range t.tops {
			t.foreach(top, f)
		}
	}
}

func (t *SimpleTree) foreach(top TreeNode, f func(TreeNode)) {
	f(top)
	children := top.GetChildren()
	if children == nil {
		return
	}

	for _, child := range children {
		t.foreach(child, f)
	}
}

// BuildTree 只是构建
func BuildTree(nodes []TreeNode) *SimpleTree {
	var resp = &SimpleTree{tops: make([]TreeNode, 0)}
	var keyMap = make(map[string]TreeNode)
	for _, node := range nodes {
		// 判断当前节点是否根节点
		if node.IsTop() {
			resp.tops = append(resp.tops, node)
		}
		keyMap[node.GetKey()] = node
	}

	for key, node := range keyMap {
		// 获取当前节点父节点的key
		pKey := node.GetParentKey()
		// 判断当前列表中是否存在该节点的父节点
		if _, ok := keyMap[pKey]; ok {
			// 两者建立绑定关系
			keyMap[pKey].AppendChild(keyMap[key])
			keyMap[key].BindingParent(keyMap[pKey])
		}
	}
	return resp
}
