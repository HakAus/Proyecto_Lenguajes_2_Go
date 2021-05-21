package trees

import (
	"fmt"
)

type AVLTree struct {
	root *Node
}

func (tree *AVLTree) Insert(toInsert int) int {
	fmt.Println("Avl Insert")
	tree.root = nil
	return 0
}

func (tree *AVLTree) Search(toFind int) (bool, int) {
	return TreeSearch(toFind)
}
