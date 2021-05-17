package trees

import "fmt"

type AvlTree struct {
	root *Node
}

func (tree *AvlTree) Insert(toInsert int) int {
	fmt.Println("Avl Insert")
	tree.root = nil
	return 0
}

func (tree *AvlTree) Search(toFind int) (bool, int) {
	return TreeSearch(toFind)
}
