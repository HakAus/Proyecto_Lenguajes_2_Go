package trees

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
	GetRoot() *Node
}

func TreeSearch(tree Tree, toFind int) (bool, int) {
	root := tree.GetRoot()
	find := false
	count := 0
	return TreeSearchRec(root, toFind, &find, &count)
}

func TreeSearchRec(node *Node, toFind int, found *bool, comparations *int) (bool, int) {
	emptyNode := Node{}
	if *node == emptyNode {
		*comparations += 1
	} else {
		if node.value == toFind {
			*found = true
			*comparations += 2
		} else {
			if toFind <= node.value {
				*comparations += 3
				TreeSearchRec(node.leftChildren, toFind, found, comparations)
			} else {
				*comparations += 3
				TreeSearchRec(node.rightChildren, toFind, found, comparations)
			}
		}
	}
	return *found, *comparations
}
