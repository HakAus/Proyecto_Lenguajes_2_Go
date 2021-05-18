package main

import (
	"fmt"

	"github.com/Xfighter32/Proyecto_Lenguajes_2_Go/trees"
)

func main() {
	fmt.Println("Hello, world!")
	//at := trees.AvlTree{}
	bt := trees.BinaryTree{Root: &trees.Node{}, CurrentKey: 0}
	var tree trees.Tree
	fmt.Print(bt.Root)
	tree = &bt
	fmt.Println(&bt.Root, "Desde main Root")
	fmt.Println(&bt, "Desde main bt")
	fmt.Println(&tree, "Desde main tree")
	fmt.Println(tree.Insert(2), "Comparaciones")
	fmt.Println(bt.Root)
	fmt.Println(bt.CurrentKey)
	fmt.Println(tree.Insert(4), "Comparaciones")
	fmt.Println(bt.Root)
	tree.Search(2)
}
