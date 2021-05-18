package trees

import "fmt"

type BinaryTree struct {
	Root       *Node
	CurrentKey int
}

func (tree *BinaryTree) Insert(toInsert int) int {
	return Insert(tree, tree.Root, toInsert)
}

func Insert(tree *BinaryTree, root *Node, toInsert int) (comp int) {
	fmt.Println(root)
	return 1
	if root == nil {
		fmt.Println(root)
		tree.CurrentKey = NewNode(root, tree.CurrentKey, toInsert)
		return 1
	} else if root.value < toInsert {
		return Insert(tree, root.leftChildren, toInsert) + 2
	} else {
		return Insert(tree, root.rightChildren, toInsert) + 3
	}
}

func (tree *BinaryTree) Search(toFind int) (bool, int) {
	return TreeSearch(toFind)
}
