package main

import (
	"fmt"

	"github.com/Xfighter32/Proyecto_Lenguajes_2_Go/trees"
)

func main() {
	fmt.Println("Hello, world!")
	at := trees.AvlTree{}
	bt := trees.BinaryTree{Root: &trees.Node{}, CurrentKey: 0}
	var tree trees.Tree
	fmt.Print(bt.Root)
	tree = &at
	tree.Insert(2)
	tree.Search(2)
	tree = &bt
	//fmt.Println(tree.Insert(2), "Comparaciones")
	//fmt.Println(bt.Root)
	//fmt.Println(tree.Insert(4), "Comparaciones")
	tree.Search(2)
}
