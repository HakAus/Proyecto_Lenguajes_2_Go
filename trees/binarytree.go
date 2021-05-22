package trees

import "fmt"

type BinaryTree struct {
	key           int
	value         int
	leftChildren  *BinaryTree
	rightChildren *BinaryTree
}

// INTERFACE

func (tree *BinaryTree) Insert(keyToInsert int) int {
	if tree.IsEmpty() {
		*tree = NewBinaryTree(keyToInsert, 1)
		return +1
	} else {
		if keyToInsert == tree.key {
			tree.value++
			return +1
		} else if keyToInsert < tree.key {
			return tree.leftChildren.Insert(keyToInsert) + 2
		} else {
			return tree.rightChildren.Insert(keyToInsert) + 2
		}
	}
}

func (tree *BinaryTree) Search(toFind int) (bool, int) {
	var itree ITree = tree
	return TreeSearch(itree, toFind)
}

func (tree *BinaryTree) IsEmpty() bool {
	return tree.key == 0 && tree.value == 0 && tree.leftChildren == nil && tree.rightChildren == nil
}

func (tree *BinaryTree) IsBranchless() bool {
	return tree.leftChildren == nil && tree.rightChildren == nil
}

func (tree *BinaryTree) GetKey() int {
	return tree.key
}

func (tree *BinaryTree) GetValue() int {
	return tree.value
}

func (tree *BinaryTree) GetLeftChildren() ITree {
	return tree.leftChildren
}

func (tree *BinaryTree) GetRightChildren() ITree {
	return tree.rightChildren
}

func (tree *BinaryTree) IRD() {
	Binary_IRD_Rec(tree)
}

// LOCAL
func NewBinaryTree(newKey int, valueToInsert int) BinaryTree {
	newTree := BinaryTree{
		key:           newKey,
		value:         valueToInsert,
		leftChildren:  &BinaryTree{},
		rightChildren: &BinaryTree{}}
	return newTree
}

func Binary_IRD_Rec(tree *BinaryTree) {
	if !tree.IsBranchless() {
		Binary_IRD_Rec(tree.leftChildren)
		fmt.Println("key: ", tree.key, "value: ", tree.value)
		Binary_IRD_Rec(tree.rightChildren)
	}
}

func (root *BinaryTree) rotateLeft() {
	var newRoot *BinaryTree = root.rightChildren
	root.rightChildren = newRoot.leftChildren
	var newLeft BinaryTree = *root
	newRoot.leftChildren = &newLeft
	*root = *newRoot
	newRoot = nil
}

func (root *BinaryTree) rotateRight() {
	var newRoot *BinaryTree = root.leftChildren
	root.leftChildren = newRoot.rightChildren
	var newRight BinaryTree = *root
	newRoot.rightChildren = &newRight
	*root = *newRoot
	newRoot = nil
}

//Toma un arbol y lo convierte en una linaea rotando los nodos hacia la derecha
func (root *BinaryTree) backboneTree() bool {
	fmt.Println(root.key)
	if root.IsEmpty() {

		return false
	}
	if root.leftChildren.IsEmpty() {
		return root.rightChildren.backboneTree()
	} else {
		root.rotateRight()
		root.backboneTree()
	}
	return false
}

//Calcular cual sera la altura del arbol y repetir el proceso el numero de niveles completos
//Se hace una rotacion inicial. Donde los nodos del nivel inferios son calculador y se rotan en la primera fase
//Toma el backbone creado y lo convierte en un arbol rotando a la izquierda el resto de nodos impares

func (root *BinaryTree) treeBackbone() {
	//Altura deseada con log2(totalDeNodos).ceil()
	desiredNodes := 2
	//desiredNodes -= 1
	fmt.Println("Primera parte de fila a arbol")
	root.treeBackboneFersto(&desiredNodes)
	fmt.Println("Segunda parte de fila a arbol")
	root.treeBackboneSecando()
}

func (root *BinaryTree) treeBackboneFersto(currentNodes *int) bool {
	if *currentNodes == 0 { //No es necesario verificar los nodos vacios. Porque la cantidad de nodos deseada evita ese tpo de errores.
		return false
	} else {
		*currentNodes -= 1
		root.rightChildren.rightChildren.treeBackboneFersto(currentNodes)
		root.rotateLeft()
	}
	return false
}

func (root *BinaryTree) treeBackboneSecando() bool {
	if root.rightChildren.IsEmpty() { //No se si sean necesarias las dos comparaciones. Para mantener la funcionalidad en casos donde no hay hijos derechos
		return false
	} else if root.rightChildren.rightChildren.IsEmpty() { //Se avanza de dos nodos a la vez para solamente alterar los impares
		return false
	} else {
		fmt.Println("Secando", root.key)
		root.rightChildren.rightChildren.treeBackboneSecando() //Puede quitarse el return?
		fmt.Println("AfterSecando", root.key)
		root.rotateLeft()
	}
	fmt.Println("Final", root.key)
	return false
}

func (root *BinaryTree) BalanceDsw() {
	fmt.Println("Start tree to backbone")
	root.backboneTree()
	root.PrintToRight()
	fmt.Println("Start backbone to tree")
	root.treeBackbone()
	root.rotateLeft()
}

func (root *BinaryTree) PrintToRight() bool {
	fmt.Println("ToRight")
	if root.IsEmpty() {
		return false
	} else {
		fmt.Println(root.key)
		fmt.Println(root.leftChildren)
		return root.rightChildren.PrintToRight()
	}
}

func (avl *BinaryTree) ToString(count int) string {
	var indent string;
	count++
	if !avl.IsEmpty() {
		var buff string;
		buff += "[Key: " + fmt.Sprintf("%d", avl.key)
		buff += "| Value: " + fmt.Sprintf("%d",avl.value)
		for i := 0; i < count; i++ {
			indent += "\t"
		}
		buff += "\n" + indent + avl.leftChildren.ToString(count)
		buff += "\n" + indent + avl.rightChildren.ToString(count)
		return buff
	}
	return "Empty"
}