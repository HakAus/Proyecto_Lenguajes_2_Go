package main

import (
	"fmt"

	"github.com/Xfighter32/Proyecto_Lenguajes_2_Go/randgen"
	"github.com/Xfighter32/Proyecto_Lenguajes_2_Go/trees"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(randgen.GetRandomArray(17,200))
	
	bt := trees.BinaryTree{}
	
	var tree trees.ITree	// I can hold values of interface type "tree"
	// fmt.Print(bt.Root)
	fmt.Println("\nlonely interface: ", tree)
	tree = &bt
	fmt.Println("\nbinary tree: ", tree)
	tree.Insert(10)
	tree.Insert(2)
	tree.Insert(2)
	tree.Insert(4)
	tree.Insert(3)
	tree.Insert(50)
	tree.Insert(7)
	tree.Insert(8)
	tree.IRD()
	// fmt.Println(bt.Root, "Desde main Root")
	// fmt.Println(bt, "Desde main bt")
	// fmt.Println(tree, "Desde main tree")
	// fmt.Println(tree.Insert(2), "Comparaciones")
	// fmt.Println(bt.Root)
	// fmt.Println(bt.CurrentKey)
	// fmt.Println(tree.Insert(4), "Comparaciones")
	// fmt.Println(bt.Root)
	// fmt.Println(tree.Insert(10))
	// fmt.Println(tree.Insert(1))
	// fmt.Println(tree.Search(10))
	fmt.Println(tree)
	
}
