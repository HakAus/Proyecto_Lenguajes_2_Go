package trees

import (
	"fmt"
	"math"
)

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
			return +2
		} else if keyToInsert < tree.key {
			return tree.leftChildren.Insert(keyToInsert) + 3
		} else {
			return tree.rightChildren.Insert(keyToInsert) + 3
		}
	}
}

func (tree *BinaryTree) Search(toFind int) (bool, int) {
	var itree ITree = tree
	return TreeSearch(itree, toFind)
}

func (tree *BinaryTree) IsEmpty() bool {
	// Todo nuevo nodo es insertado con value 1 por tanto con verificar que sea 0 es suficiente.
	return tree.value == 0
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

func (tree *BinaryTree) IRD() []int {
	var answer []int
	Binary_IRD_Rec(tree, &answer)
	return answer
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

func Binary_IRD_Rec(tree *BinaryTree, answer *[]int) {
	if !tree.IsBranchless() {
		Binary_IRD_Rec(tree.leftChildren, answer)
		*answer = append(*answer, tree.key)
		Binary_IRD_Rec(tree.rightChildren, answer)
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

//Toma un arbol y lo convierte en una linea rotando los nodos hacia la derecha
func (root *BinaryTree) backboneTree() {
	if !root.IsEmpty() {
		if root.leftChildren.IsEmpty() { 
			root.rightChildren.backboneTree()
		} else { 
			root.rotateRight()
			root.backboneTree() 
		}
	}
}

/* Calcular cual sera la altura del arbol y repetir el proceso el numero de niveles completos
Se hace una rotacion inicial. Donde los nodos del nivel inferios son calculados y se rotan en la primera fase
Toma el backbone creado y lo convierte en un arbol rotando a la izquierda el resto de nodos impares*/
func (root *BinaryTree) treeBackbone() {
	totalNodes := CountNodes(root)
	balancedLevels := math.Floor(math.Log2(float64(totalNodes)))
	desiredNodes := totalNodes - int(math.Pow(2, float64(balancedLevels))) + 1
	times := totalNodes
	root.treeBackboneFersto(&desiredNodes)
	for times > 2 {
		times /= 2
		root.treeBackboneSecando()
	}
}

//Rotaciones y punteros con algo de antes
func (root *BinaryTree) treeBackboneFersto(currentNodes *int) {
	//No es necesario verificar los nodos vacios. Porque la cantidad de nodos deseada evita ese tipo de errores.
	if *currentNodes != 0 {
		*currentNodes -= 1
		root.rightChildren.rightChildren.treeBackboneFersto(currentNodes)
		root.rotateLeft()
	}
}

func (root *BinaryTree) treeBackboneSecando() {
	if root.rightChildren.IsEmpty() {
		return
	} else if root.rightChildren.rightChildren.IsEmpty() { //Se avanza de dos nodos a la vez para solamente alterar los impares
		return
	} else {
		root.rightChildren.rightChildren.treeBackboneSecando()
		root.rotateLeft()
	}
}

func (root *BinaryTree) BalanceDsw() {
	root.backboneTree()
	root.treeBackbone()
}

func (root *BinaryTree) PrintToRight() bool {
	if root.IsEmpty() {
		return false
	} else {
		fmt.Print(root.key, ",")
		return root.rightChildren.PrintToRight()
	}
}

func (avl *BinaryTree) ToString(indentAmount int) string {
	var indent string
	indentAmount++
	if !avl.IsEmpty() {
		var buff string
		buff += "[Key: " + fmt.Sprintf("%d", avl.key)
		buff += "| Value: " + fmt.Sprintf("%d", avl.value)
		for i := 0; i < indentAmount; i++ {
			indent += "\t"
		}
		buff += "\n" + indent + avl.leftChildren.ToString(indentAmount)
		buff += "\n" + indent + avl.rightChildren.ToString(indentAmount)
		return buff
	}
	return "Empty"
}
