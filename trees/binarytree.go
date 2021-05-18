package trees

import "fmt"

type BinaryTree struct {
	Root       *Node
	CurrentKey int
}

func (tree *BinaryTree) Insert(toInsert int) int {
	fmt.Println("Start")
	fmt.Println("Desde Insert Root", &tree.Root, " ", tree.Root)
	fmt.Println("Desde Insert Tree ", tree)
	return tree.Root.Insert(tree, toInsert)
}

func (node *Node) Insert(tree *BinaryTree, toInsert int) (comp int) {
	emptyNode := Node{}
	if *node == emptyNode {
		*node, tree.CurrentKey = NewNode(tree.CurrentKey, toInsert)
		return 1
	} else if node.value <= toInsert {
		fmt.Println("OwO")
		return node.leftChildren.Insert(tree, toInsert) + 2
	} else {
		return node.rightChildren.Insert(tree, toInsert) + 2
	}
}

func (tree *BinaryTree) Search(toFind int) (bool, int) {
	return TreeSearch(toFind)
}
