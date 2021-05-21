package trees

import (
	"fmt"
)


type AVlTree struct {
	Root *Node
}

func (tree *AVlTree) Insert(toInsert int) int {
	fmt.Println("Avl Insert")
	return 0
}

func (tree *AVlTree) GetRoot() *Node{
	return tree.Root;
}

func (tree *AVlTree) Search(toFind int) (bool, int) {
	var itree Tree = tree
	return TreeSearch(itree,toFind)
}
