package main

import (
	// "github.com/Xfighter32/Proyecto_Lenguajes_2_Go/randgen"
	//"fmt"

	"github.com/Xfighter32/Proyecto_Lenguajes_2_Go/trees"
)

func main() {
	// fmt.Println(randgen.GetRandomArray(17,200))

	//avl := trees.AVLTree{}
	//var tree trees.ITree = &avl
	bt := trees.BinaryTree{}
	var tree trees.ITree = &bt

	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(20)
	tree.Insert(7)
	tree.Insert(15)
	tree.Insert(30)
	tree.Insert(25)
	tree.Insert(40)
	tree.Insert(23)
	bt.BalanceDsw()
	// tree.IRD()
	//fmt.Println("\n", bt.PrintToRight())
	//fmt.Println(bt.ToString(0))
}
