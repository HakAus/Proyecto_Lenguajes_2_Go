package trees

import "fmt"

type Node struct {
	key           int
	value         int
	leftChildren  *Node
	rightChildren *Node
}

type AVLNode struct {
	
}

func NewNode(newKey int, toInsert int) (Node, int) {
	newNode := Node{
		key:           newKey,
		value:         toInsert,
		leftChildren:  &Node{},
		rightChildren: &Node{}}
	return newNode, newKey + 1
}

type Tree interface {
	Insert(int) int
	Search(int) (bool, int)
}

func TreeSearch(int) (bool, int) {
	fmt.Println("Same Search")
	return false, 0
}
