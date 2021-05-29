package trees

import (
	"fmt"

	"github.com/Xfighter32/Proyecto_Lenguajes_2_Go/randgen"
)

type TreeStats struct {
	tree        ITree
	comparisons int
	size        int
	isDSW       bool
}

func StartStats() {
	// Casos de prueba
	sizes := []int{200, 400, 600, 800, 1000}

	// Llenado de árboles
	for _, size := range sizes {
		fmt.Printf("-----------------------------------------------%d-----------------------------------------------\n", size)
		bt := TreeStats{tree: &BinaryTree{}, size: size, isDSW: false}
		dsw := TreeStats{tree: &BinaryTree{}, size: size, isDSW: true}
		avl := TreeStats{tree: &AVLTree{}, size: size, isDSW: false}
		bt.populate(size)
		dsw.populate(size)
		dsw.TransformDSW()
		avl.populate(size)
		fmt.Println(dsw.tree.ToString(0))

		// Se genera una secuencia de 10 000 números aleatorios
		biglist := randgen.GetRandomArray(17, 10000,200)

		// Se realizan las busquedas en cada uno de los árboles
		bt.search(biglist)
		dsw.search(biglist)
		avl.search(biglist)

		// Se hacen las estadísticas
		bt.getStats()
		dsw.getStats()
		avl.getStats()
		fmt.Println("------------------------------------------------------------------------------------------------")
	}

	fmt.Println("Finished!")

}

func (ts TreeStats) getInfo() {
	switch ts.tree.(type) {
	case *BinaryTree:
		if ts.isDSW {
			fmt.Printf("\n **** Estadísticas de árbol binario transormado por DSW | Arreglo: %d **** \n \n", ts.size)
		} else {
			fmt.Printf("\n **** Estadísticas de árbol binario | Arreglo: %d ****\n \n", ts.size)
		}
	case *AVLTree:
		fmt.Printf("\n **** Estadísticas de árbol AVL | Arreglo: %d ****\n \n", ts.size)
	}
}

func (ts TreeStats) getStats() {
	ts.getInfo()
	// Calcular la altura
	fmt.Printf("Altura máxima: %d \n", Height(ts.tree))
	// Altura promedio
	fmt.Printf("Altura promedio: %f \n", GetAvgHeight(ts.tree))
	// Densidad del árbol
	fmt.Printf("Densidad del árbol: %f \n", Density(ts.tree))
	// Cantidad total de comparaciones realizadas
	fmt.Printf("Cantidad total de comparaciones: %d \n", ts.comparisons)
	// Cantidad promedio de comparaciones para las inserciones realizadas
	fmt.Printf("Cantidad promedio de comparaciones: %f \n", GetAvgComparisons(ts.tree))
}

//Utiliza el arbol dentro de tree_stats y le inserta todos los elementos del array
func (ts TreeStats) populate(size int) {
	randnums := randgen.GetRandomArray(17, size, 200)
	for _, value := range randnums {
		ts.tree.Insert(value)
	}
}

func (ts TreeStats) TransformDSW() {
	switch typedTree := ts.tree.(type) {
	case *BinaryTree:
		typedTree.BalanceDsw()
	default:
	}
}

func (ts *TreeStats) search(slice []int) {
	for _, value := range slice {
		_, comp := ts.tree.Search(value)
		ts.comparisons += comp
	}
}
