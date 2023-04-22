package treex

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	var list = []TreeNode{
		NewDefaultTree("01", "", "01_val"),
		NewDefaultTree("0101", "01", "01_val"),
		NewDefaultTree("0102", "01", "01_val"),
		NewDefaultTree("0103", "01", "01_val"),
		NewDefaultTree("0104", "01", "01_val"),
		NewDefaultTree("0105", "01", "01_val"),
	}

	tree := BuildTree(list)
	tree.Foreach(func(node TreeNode) {
		fmt.Println("key -> ", node.GetKey())
	})
}
