package trees

import (
	"fmt"
)


type AvlTree struct {
	Root *Node
}

func (tree *AvlTree) Insert(toInsert int) int {
	fmt.Println("Avl Insert")
	return 0
}

func (tree *AvlTree) GetRoot() *Node{
	return tree.Root;
}

func (tree *AvlTree) Search(toFind int) (bool, int) {
	var itree Tree = tree
	return TreeSearch(itree,toFind)
}
