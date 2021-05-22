package trees

import (
	"fmt"
)

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
	change := false
	fmt.Println("\nStart")
	fmt.Printf("original avl address: %p \n", tree)
	AVL_Insert(tree, keyToInsert,&comparisons,&change)
	fmt.Println("Luego de la insercion")
	// fmt.Println(tree.toString(0))
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

func AVL_Insert(avl *AVLTree, keyToInsert int, comparisons *int, change *bool) *AVLTree {
	fmt.Printf("avl address: %p | avl key: %d | key to insert: %d \n",avl, avl.key, keyToInsert)
	if avl.IsEmpty() {
		*avl = NewAVLTree(keyToInsert, 1)
		fmt.Printf("New tree added: %p | new key: %d \n", avl, avl.key)
		*change = true
	} else {
		if keyToInsert == avl.key {
			avl.value++
			*comparisons += 1
		} else if keyToInsert < avl.key {
			*comparisons += 2
			fmt.Println("Going left ...")
			avl.leftChildren = AVL_Insert(avl.leftChildren, keyToInsert, comparisons, change)
			fmt.Printf("Unwound avl address: %p | key: %d \n", avl, avl.key)

			if (*change) {
				switch avl.bf {
				case rightImbalanced:
					// fmt.Println("rightImbalanced")
					avl.bf = balanced
					*change = false
				case balanced:
					// fmt.Println("balanced")
					avl.bf = leftImbalanced
				case leftImbalanced:
					// fmt.Println("leftImbalanced")
					n1 := avl.leftChildren;
					if n1.bf == leftImbalanced {
						avl.rotateRight()
					} else {
						avl.rotateLeftRight()
					}
					*change = false
				}
				fmt.Println("Luego del balanceo")
			}
		} else {
			*comparisons += 2
			fmt.Println("Going right ...")
			avl.rightChildren = AVL_Insert(avl.rightChildren, keyToInsert, comparisons, change)
			fmt.Printf("Unwound avl address: %p | key: %d \n", avl, avl.key)
			
			if (*change) {
				switch avl.bf {
				case rightImbalanced:
					// fmt.Println("rightImbalanced")
					n1 := avl.rightChildren
					if n1.bf == rightImbalanced {
						avl.rotateLeft()
						fmt.Printf("FUR REAL: %p", avl)
					} else {
						avl.rotateRightLeft()
					}
					*change = false
				case balanced:
					// fmt.Println("balanced")
					avl.bf = rightImbalanced
				case leftImbalanced:
					// fmt.Println("leftImbalanced")
					avl.bf = balanced
					*change = false
				}
			}
			fmt.Println("Luego del balanceo")
		}
	}
	fmt.Println("BEFORE RETURNING")
	fmt.Printf("avl address: %p | avl key: %d \n", avl, avl.key)
	fmt.Printf("avl left address: %p | avl key: %d \n", avl.leftChildren, avl.leftChildren.key)
	fmt.Printf("avl right address: %p | avl key: %d \n", avl.rightChildren, avl.rightChildren.key)
	return avl
}

func (root *AVLTree) rotateLeft() {
	var newRoot *AVLTree = root.rightChildren
	root.rightChildren = newRoot.leftChildren
	var newLeft AVLTree = *root
	newRoot.leftChildren = &newLeft
	*root = *newRoot
	newRoot = nil
	oldRoot := root.leftChildren

	// Adjust the balance factor
	if root.bf == rightImbalanced {
		oldRoot.bf = balanced
		root.bf = balanced
	} else {
		oldRoot.bf = rightImbalanced
		root.bf = leftImbalanced
	}
}
func (root *AVLTree) rotateRight() {
	var newRoot *AVLTree = root.leftChildren
	root.leftChildren = newRoot.rightChildren
	var newRight AVLTree = *root
	newRoot.rightChildren = &newRight
	*root = *newRoot
	newRoot = nil
	oldRoot := root.rightChildren

	// Adjust the balance factor
	if root.bf == leftImbalanced {
		oldRoot.bf = balanced
		root.bf = balanced
	} else {
		oldRoot.bf = leftImbalanced
		root.bf = rightImbalanced
	}
}

func (root *AVLTree) rotateLeftRight() {
	var newLeft *AVLTree = root.leftChildren
	var newRoot *AVLTree = newLeft.rightChildren
	newLeft.rightChildren = newRoot.leftChildren
	newRoot.leftChildren = newLeft
	root.leftChildren = newRoot.rightChildren
	var newRight AVLTree = *root
	newRoot.rightChildren = &newRight
	*root = *newRoot
	oldRoot := root.rightChildren
	oldRootLeft := root.leftChildren

	// Adjust balance factor
	if root.bf == rightImbalanced {
		oldRootLeft.bf = leftImbalanced
	} else {
		oldRootLeft.bf = balanced
	}

	if root.bf == leftImbalanced {
		oldRoot.bf = rightImbalanced
	} else {
		oldRoot.bf = balanced
	}

	root.bf = balanced
}
func (root *AVLTree) rotateRightLeft() {

	var newRight *AVLTree = root.rightChildren
	var newRoot *AVLTree = newRight.leftChildren
	newRight.leftChildren = newRoot.rightChildren
	newRoot.rightChildren = newRight
	root.rightChildren = newRoot.leftChildren
	var newLeft AVLTree = *root
	newRoot.leftChildren = &newLeft
	*root = *newRoot
	oldRoot := root.leftChildren
	oldRootRight := root.rightChildren

	// Adjust balance factor
	if root.bf == rightImbalanced {
		oldRoot.bf = leftImbalanced
	} else {
		oldRoot.bf = balanced
	}

	if root.bf == leftImbalanced {
		oldRootRight.bf = rightImbalanced
	} else {
		oldRootRight.bf = balanced
	}
	root.bf = balanced
}


func (avl *AVLTree) ToString(count int) string {
	var indent string;
	count++
	if !avl.IsEmpty() {
		var buff string;
		buff += "[Key: " + fmt.Sprintf("%d", avl.key)
		buff += "| Value: " + fmt.Sprintf("%d",avl.value)
		buff += "| BF: " + fmt.Sprintf("%d",avl.bf) + "]"
		for i := 0; i < count; i++ {
			indent += "\t"
		}
		buff += "\n" + indent + avl.leftChildren.ToString(count)
		buff += "\n" + indent + avl.rightChildren.ToString(count)
		return buff
	}
	return "Empty"
}