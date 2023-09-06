package tree

import "fmt"

type Node interface {
	Key() string       // 当前key
	ParentKey() string // 父级key
	Append(...Node)    // 添加子节点
}

type Nodes []Node

type Tree struct {
	topKey string
	nodes  []Node
}

func (t Tree) Nodes() []Node {
	return t.nodes
}

func BuildTree(topKey string, nodes ...Node) *Tree {
	var tree = &Tree{
		topKey: topKey,
		nodes:  make([]Node, 0),
	}
	if len(nodes) == 0 {
		return tree
	}

	var indexMap = make(map[string][]int)
	var topNodes = make([]Node, 0)

	for i := 0; i < len(nodes); i++ {
		var now = nodes[i]
		var index = fmt.Sprintf("%s_%s", now.ParentKey(), now.Key())
		if v, ok := indexMap[index]; ok {
			indexMap[index] = append(v, i)
		} else {
			indexMap[index] = []int{i}
		}
		// 是否是根节点
		if topKey == now.ParentKey() {
			topNodes = append(topNodes, now)
			continue
		}
	}
	// 空的父节点
	if len(topNodes) == 0 || len(indexMap) == 0 {
		tree.nodes = topNodes
		return tree
	}
	for i := 0; i < len(topNodes); i++ {
		build(topNodes[i], indexMap, nodes)
	}

	tree.nodes = topNodes
	return tree
}

func build(now Node, relation map[string][]int, nodes []Node) {
	index := fmt.Sprintf("%s_%s", now.ParentKey(), now.Key())
	if v, ok := relation[index]; ok {
		for i := 0; i < len(v); i++ {
			var node = nodes[v[i]]
			build(node, relation, nodes)
			now.Append(node)
		}
	}
}
