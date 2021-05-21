package trees

import "fmt"

type BinaryTree struct {
	key			   int
	value          int
	leftChildren   *BinaryTree
	rightChildren  *BinaryTree
}

// INTERFACE

func (tree *BinaryTree) Insert(keyToInsert int) int {

	if tree.IsBranchless() {
		if keyToInsert == tree.key {
			tree.value++
			return + 1
		} else {
			*tree = NewBinaryTree(keyToInsert, 1)
			return + 1
		}
	}
	if keyToInsert == tree.key {
		tree.value++
		return + 1
	} else if keyToInsert < tree.key {
		return tree.leftChildren.Insert(keyToInsert) + 2
	} else {
		return tree.rightChildren.Insert(keyToInsert) + 2
	}
}

func (tree *BinaryTree) Search(toFind int) (bool, int) {
	var itree ITree = tree
	return TreeSearch(itree, toFind)
}

func (tree *BinaryTree) IsEmpty() bool {
	return tree.key == 0 && tree.value == 0 && tree.leftChildren == nil && tree.rightChildren == nil
}

func (tree *BinaryTree) IsBranchless() bool {
	return tree.leftChildren == nil && tree.rightChildren == nil
}

func (tree *BinaryTree) GetKey() int {
	return tree.key
}

func (tree *BinaryTree) GetValue() int {
	return tree.value
}

func (tree *BinaryTree) GetLeftChildren() ITree {
	return tree.leftChildren
}

func (tree * BinaryTree) GetRightChildren() ITree {
	return tree.rightChildren
}

func (tree *BinaryTree) IRD() {
	Binary_IRD_Rec(tree)
}

// LOCAL
func NewBinaryTree(newKey int, valueToInsert int) (BinaryTree) {
	newTree := BinaryTree{
		key:           newKey,
		value:         valueToInsert,
		leftChildren:  &BinaryTree{},
		rightChildren: &BinaryTree{}}
	return newTree
}

func Binary_IRD_Rec(tree *BinaryTree) {
	if !tree.IsBranchless() {
		Binary_IRD_Rec(tree.leftChildren)
		fmt.Println("key: ",tree.key, "value: ", tree.value)
		Binary_IRD_Rec(tree.rightChildren)
	}
}