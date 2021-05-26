package trees

import (
	"fmt"
	"testing"
)

// Expected tree
//    2
//      10
//     4
func TestTreeHeight3(t *testing.T) {
	bt := BinaryTree{}
	bt.Insert(2)
	bt.Insert(10)
	bt.Insert(4)
	bt.Insert(2)

	var tree ITree = &bt
	got := Height(tree)
	if got != 3 {
		t.Errorf("Height(bt) = %d ; want 3", got)
	}
}

func TestTreeHeight6(t *testing.T) {
	var values = []int{10, 25, 190, 150, 192, 29, 20, 42, 12}
	bt := BinaryTree{}
	var tree ITree = &bt
	for _, value := range values {
		tree.Insert(value)
	}

	got := Height(tree)
	if got != 6 {
		t.Errorf("Height(bt) = %d ; want 6", got)
	}
}

func TestDSW(t *testing.T) {
	var values = []int{10, 5, 20, 7, 15, 30, 25, 40, 23, 5}
	var orderedValues = []int{5, 7, 10, 15, 20, 23, 25, 30, 40}
	bt := BinaryTree{}
	var tree ITree = &bt
	for _, value := range values {
		tree.Insert(value)
	}

	bt2 := BinaryTree{}
	var tree2 ITree = &bt2
	tree2.Insert(23)
	tree2.Insert(15)
	tree2.Insert(30)
	tree2.Insert(7)
	tree2.Insert(20)
	tree2.Insert(25)
	tree2.Insert(40)
	tree2.Insert(5)
	tree2.Insert(5)
	tree2.Insert(10)

	bt.BalanceDsw()
	fmt.Println(bt.ToString(0))
	got := Height(tree)
	if got != 4 {
		t.Errorf("Height(bt) = %d ; want 4", got)
	}

	var got1 = tree.IRD()
	if !EqualSlices(got1, orderedValues) {
		t.Errorf("tree.IRD() = %v ; want %v", got1, orderedValues)
	}

	got2 := CompareTrees(tree, tree2)
	if !got2 {
		t.Error("CompareTrees(tree,tree2) dont match")
	}

	got3 := Height(tree) == Height(tree2)
	if !got3 {
		t.Error("Tree heights don't match")
	}
}

// To compare slices
func EqualSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

//			30
//		  23		 40
//	 15 	25
//	7 20
//5 10
