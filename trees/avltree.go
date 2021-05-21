package trees

import "fmt"

// import (
// 	"fmt"
// )
type BalanceFactor int 

const (
	leftBalanced BalanceFactor = -1
	balanced BalanceFactor = 0
	rightBalanced BalanceFactor = 1
	leftImbalanced = -2
	rightImbalanced = 2
)

type AVlTree struct {
	key	int
	value int
	bf BalanceFactor
	leftChildren *AVlTree
	rightChildren *AVlTree
}

// INTERFACE

func (tree *AVlTree) Insert(keyToInsert int) int {
	if tree.IsBranchless() {
		if keyToInsert == tree.key {
			tree.value++
			return + 1
		} else {
			*tree = NewAVLTree(keyToInsert, 1)
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

func (tree *AVlTree) Search(toFind int) (bool, int) {
	var itree ITree = tree
	return TreeSearch(itree, toFind)
}

func (tree *AVlTree) IsEmpty() bool {
	return tree.key == 0 && tree.value == 0 && tree.leftChildren == nil && tree.rightChildren == nil
}

func (tree *AVlTree) IsBranchless() bool {
	return tree.leftChildren == nil && tree.rightChildren == nil
}

func (tree *AVlTree) GetKey() int {
	return tree.key
}

func (tree *AVlTree) GetValue() int {
	return tree.value
}

func (tree *AVlTree) GetLeftChildren() ITree {
	return tree.leftChildren
}

func (tree * AVlTree) GetRightChildren() ITree {
	return tree.rightChildren
}

func (tree *AVlTree) IRD() {
	AVL_IRD_Rec(tree)
}

// LOCAL
func NewAVLTree(newKey int, valueToInsert int) (AVlTree) {
	newTree := AVlTree{
		key:           newKey,
		value:         valueToInsert,
		bf:			   balanced,
		leftChildren:  &AVlTree{},
		rightChildren: &AVlTree{}}
	return newTree
}

func AVL_IRD_Rec(tree *AVlTree) {
	if !tree.IsBranchless() {
		AVL_IRD_Rec(tree.leftChildren)
		fmt.Println("key: ",tree.key, "value: ", tree.value)
		AVL_IRD_Rec(tree.rightChildren)
	}
}