package main

import (
	"fmt"

	"github.com/Xfighter32/Proyecto_Lenguajes_2_Go/trees"
)

func main() {
	fmt.Println("Hello, world!")

	var at = &trees.AvlTree{}
	fmt.Println("Hello, world1!")
	var bt = &trees.BinaryTree{}
	fmt.Println("Hello, world2!")
	var tree trees.Tree
	fmt.Println("Hello, xd!")
	tree = at
	tree.Insert(2)
	tree.Search(2)
	tree = bt
	tree.Insert(2)
	tree.Search(2)
}
