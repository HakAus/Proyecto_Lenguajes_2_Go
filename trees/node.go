package trees

import "fmt"

type Node struct {
	key           int
	value         int
	leftChildren  *Node
	rightChildren *Node
	weight        int8
}

func NewNode(root *Node, newKey int, toInsert int) int {
	fmt.Print(root)
	root.key = newKey
	root.value = toInsert
	root.leftChildren = &Node{}
	root.rightChildren = &Node{}
	root.weight = 0
	return newKey + 1
}

type Tree interface {
	Insert(int) int
	Search(int) (bool, int)
}

func TreeSearch(int) (bool, int) {
	fmt.Println("Same Search")
	return false, 0
}
