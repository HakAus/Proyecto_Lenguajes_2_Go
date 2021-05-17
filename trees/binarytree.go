package trees

import "fmt"

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree) Insert(toInsert int) int {
	fmt.Println("Binary Insert")
	tree.root = &Node{}
	tree.root.key = 1
	tree.root.value = 1
	tree.root.parent = nil
	tree.root.leftChildren = nil
	tree.root.rightChildren = nil
	tree.root.weight = 1
	return 0
}
func (tree *BinaryTree) Search(toFind int) (bool, int) {
	return TreeSearch(toFind)
}
