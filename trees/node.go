package trees

import "fmt"

type Node struct {
	key           int
	value         int
	parent        *Node
	leftChildren  *Node
	rightChildren *Node
	weight        int8
}

type Tree interface {
	Insert(int) int
	Search(int) (bool, int)
}

func TreeSearch(int) (bool, int) {
	fmt.Println("Same Search")
	return false, 0
}
