package trees

import (
	"fmt"
	// "math"
	"testing"

	"github.com/Xfighter32/Proyecto_Lenguajes_2_Go/randgen"
)

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

func mapNumber(value, start1, stop1, start2, stop2 int) int {
	return  int(float64((value - start1)) / float64((stop1 - start1)) * float64((stop2 - start2)) + float64(start2))
}

func BenchmarkRandomArray200(b *testing.B) {
	dict := make([]int, 200)
	nextRandom := randgen.RandomIntGenerator(23)
	fmt.Println("N: ", b.N)
	for i := 0; i < b.N; i++ {
		v := mapNumber(nextRandom(), 0, 4096, 0, 200)
		if v == 200 {
			fmt.Println("\n NOOOO ")
		}
		dict[v]++
	}
	var results string
	for _, freq := range dict {
		results += fmt.Sprintf("%d", freq) + ","
	}
	fmt.Println("RESULTADOS: ", results)
}


