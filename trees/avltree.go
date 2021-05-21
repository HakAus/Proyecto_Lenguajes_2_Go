package trees

import "fmt"

// import (
// 	"fmt"
// )
type BalanceFactor int 

const (
	leftImbalanced BalanceFactor = -1
	balanced BalanceFactor = 0
	rightImbalanced BalanceFactor = 1
)

type AVLTree struct {
	key	int
	value int
	bf BalanceFactor
	leftChildren *AVLTree
	rightChildren *AVLTree
}

// INTERFACE

func (tree *AVLTree) Insert(keyToInsert int) int {
	comparisons := 0
	tree.AVL_Insert(keyToInsert,&comparisons,false)
	return comparisons
}

func (tree *AVLTree) Search(toFind int) (bool, int) {
	var itree ITree = tree
	return TreeSearch(itree, toFind)
}

func (tree *AVLTree) IsEmpty() bool {
	return tree.key == 0 && tree.value == 0 && tree.leftChildren == nil && tree.rightChildren == nil
}

func (tree *AVLTree) IsBranchless() bool {
	return tree.leftChildren == nil && tree.rightChildren == nil
}

func (tree *AVLTree) GetKey() int {
	return tree.key
}

func (tree *AVLTree) GetValue() int {
	return tree.value
}

func (tree *AVLTree) GetLeftChildren() ITree {
	return tree.leftChildren
}

func (tree * AVLTree) GetRightChildren() ITree {
	return tree.rightChildren
}

func (tree *AVLTree) IRD() {
	AVL_IRD_Rec(tree)
}

// LOCAL
func NewAVLTree(newKey int, valueToInsert int) (AVLTree) {
	newTree := AVLTree{
		key:           newKey,
		value:         valueToInsert,
		bf:			   balanced,
		leftChildren:  &AVLTree{},
		rightChildren: &AVLTree{}}
	return newTree
}

func AVL_IRD_Rec(tree *AVLTree) {
	if !tree.IsBranchless() {
		AVL_IRD_Rec(tree.leftChildren)
		fmt.Println("key: ",tree.key, "value: ", tree.value)
		AVL_IRD_Rec(tree.rightChildren)
	}
}

func (avl *AVLTree) AVL_Insert(keyToInsert int, comparisons *int, change bool) int {
	if avl.IsEmpty() {
		*avl = NewAVLTree(keyToInsert, 1)
		change = true
	} else {
		if keyToInsert == avl.key {
			avl.value++
			*comparisons += 1
		} else if keyToInsert < avl.key {
			*comparisons += 2
			avl.leftChildren.Insert(keyToInsert)
			if (change) {
				switch avl.bf {
				case rightImbalanced:
					avl.bf = balanced
					change = false
				case balanced:
					avl.bf = leftImbalanced
				case leftImbalanced:
					n1 := avl.leftChildren;
					if n1.bf == leftImbalanced {
						avl.rotateLeftTwice()
					} else {
						avl.rotateLeftRight()
					}
					change = false
				}
			}
		} else {
			
			*comparisons += 2
			avl.rightChildren.Insert(keyToInsert)
			if (change) {
				switch avl.bf {
				case rightImbalanced:
					n1 := avl.rightChildren
					if n1.bf == rightImbalanced {
						avl.rotateRightTwice()
					} else {
						avl.rotateRightLeft()
					}
					change = false
				case balanced:
					avl.bf = rightImbalanced
				case leftImbalanced:
					avl.bf = balanced
					change = false
				}
			}
		}
	}
	return *comparisons
}

func (avl *AVLTree) rotateLeftTwice() {

}
func (avl *AVLTree) rotateRightTwice(){

}
func (avl *AVLTree) rotateLeftRight() {

}
func (avl *AVLTree) rotateRightLeft() {

}