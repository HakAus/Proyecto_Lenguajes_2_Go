package trees

type Node struct {
	key           int
	value         int
	leftChildren  *Node
	rightChildren *Node
}

func NewNode(newKey int, toInsert int) (Node, int) {
	newNode := Node{
		key:           newKey,
		value:         toInsert,
		leftChildren:  &Node{},
		rightChildren: &Node{}}
	return newNode, newKey + 1
}
