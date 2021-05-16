package trees

type Node struct{
	key int
	value int
	parent *Node 
	leftChildren *Node
	rightChildren *Node
	weight int8
}

type Tree interface{
	insert(int) (int)
	search(int) (bool,int)
}