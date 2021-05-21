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


// func zTreeSearch(tree Tree, toFind int) (bool, int) {
// 	root := &tree
// 	find := false
// 	count := 0
// 	return TreeSearchRec(root, toFind, &find, &count)
// }

// func TzreeSearchRec(node *Node, toFind int, found *bool, comparations *int) (bool, int) {
// 	emptyNode := Node{}
// 	if *node == emptyNode {
// 		*comparations += 1
// 	} else {
// 		if node.value == toFind {
// 			*found = true
// 			*comparations += 2
// 		} else {
// 			if toFind <= node.value {
// 				*comparations += 3
// 				TreeSearchRec(tree.leftChildren, toFind, found, comparations)
// 			} else {
// 				*comparations += 3
// 				TreeSearchRec(tree.rightChildren, toFind, found, comparations)
// 			}
// 		}
// 	}
// 	return *found, *comparations
// }
