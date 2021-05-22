package main

import (
	// "github.com/Xfighter32/Proyecto_Lenguajes_2_Go/randgen"
	// "fmt"

	"fmt"

	"github.com/Xfighter32/Proyecto_Lenguajes_2_Go/trees"
)

func main() {
	// fmt.Println(randgen.GetRandomArray(17,200))
	
	avl := trees.AVLTree{}
	var tree trees.ITree = &avl

	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(5)
	tree.Insert(7)
	tree.Insert(20)
	tree.Insert(13)
	tree.Insert(21)
	tree.Insert(15)
	tree.Insert(17)
	// tree.IRD()
	fmt.Println("\n", avl.ToString(0))
}
