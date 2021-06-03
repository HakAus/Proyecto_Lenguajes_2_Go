package trees

import (
	"fmt"

	"github.com/Xfighter32/Proyecto_Lenguajes_2_Go/randgen"
)

type TreeStats struct {
	tree              ITree
	comparisons       int
	insertComparisons int
	searchComparisons int
	size              int
	isDSW             bool
}

func StartStats() {
	// Casos de prueba (a)
	sizes := []int{200, 400, 600, 800, 1000}
	trees := []TreeStats{}
	for _, size := range sizes {
		fmt.Printf("-----------------------------------------------%d-----------------------------------------------\n", size)
		// Creación  de árboles vacíos para cada método (b)
		bt := TreeStats{tree: &BinaryTree{}, size: size, isDSW: false}
		dsw := TreeStats{tree: &BinaryTree{}, size: size, isDSW: true}
		avl := TreeStats{tree: &AVLTree{}, size: size, isDSW: false}

		// Inserción de todos los elementos del arreglo de tamaño "size" (b)
		btOk := bt.populate(size)
		dswOk := dsw.populate(size)
		avlOk := avl.populate(size)

		if btOk && dswOk && avlOk {
			// Transformación a DSW luego de inserción de todos los datos (b)
			dsw.TransformDSW()	

			// Se genera una secuencia de 10 000 números pseudo-aleatorios (c)
			biglist := randgen.GetRandomArray(17,10000)

			if biglist != nil {
				// Se realizan las búsquedas en cada uno de los árboles (c)
				bt.search(biglist)
				dsw.search(biglist)
				avl.search(biglist)

				// Se hacen las estadísticas (d)
				bt.getStats()
				dsw.getStats()
				avl.getStats()

				trees = append(trees, bt)
				trees = append(trees, dsw)
				trees = append(trees, avl)

				fmt.Println("------------------------------------------------------------------------------------------------")
			} else {
				fmt.Println("Hubo un problema con la lista de valores de búsqueda")
			}
		} else {
			fmt.Println("Hubo un problema poblando los árboles")
		}
	}

	fmt.Println("Terminado!")

	fmt.Println("AFTERMATH")
	getMethodStat("bt", trees, getInsertResults)
	getMethodStat("dsw", trees, getInsertResults)
	getMethodStat("avl", trees, getInsertResults)
	getMethodStat("bt", trees, getSearchResults)
	getMethodStat("dsw", trees, getSearchResults)
	getMethodStat("avl", trees, getSearchResults)
}

func getInsertResults(ts TreeStats) {
	fmt.Printf("Resultados de comparaciones de inserción en árbol de tamaño %d: %d\n", ts.size, ts.insertComparisons)
}

func getSearchResults(ts TreeStats){
    fmt.Printf("Resultados de comparaciones de búsqueda en árbol de tamaño %d: %d\n", ts.size, ts.searchComparisons)
}

func getMethodStat(treeType string, trees []TreeStats, statFunc func(TreeStats)) {
	var start int
	increase := 3

	switch treeType {
	case "bt":
		start = 0
	case "dsw":
		start = 1
	case "avl":
		start = 2
	}
	fmt.Println("Árbol ", treeType)
	for start < len(trees) {
		statFunc(trees[start])
		start += increase
	}
}

// Imprime un encabezado para las estadísticas de cada método
func (ts TreeStats) getStatsHeader() {
	switch ts.tree.(type) {
	case *BinaryTree:
		if ts.isDSW {
			fmt.Printf("\n **** DSW | Arreglo: %d **** \n \n", ts.size)
		} else {
			fmt.Printf("\n **** ABB | Arreglo: %d ****\n \n", ts.size)
		}
	case *AVLTree:
		fmt.Printf("\n **** AVL | Arreglo: %d ****\n \n", ts.size)
	}
}

// Generación de estadísticas para cada método desarrollado (d)
func (ts TreeStats) getStats() {
	// Encabezado
	ts.getStatsHeader()
	// Altura máxima del árbol
	fmt.Printf("Altura máxima: %d \n", Height(ts.tree))
	// Profundiad promedio del árbol
	fmt.Printf("Altura o profundidad promedio: %f \n", GetAvgHeight(ts.tree))
	// Densidad del árbol
	fmt.Printf("Densidad del árbol: %f \n", Density(ts.tree))
	// Cantidad total de comparaciones realizadas (búsquedas + inserciones)
	fmt.Printf("Cantidad total de comparaciones: %d \n", ts.comparisons)
	fmt.Printf("Cantidad total de comparaciones en BÚSQUEDA: %d \n", ts.searchComparisons)
	fmt.Printf("Cantidad total de comparaciones en INSERCIÓN: %d \n", ts.insertComparisons)
	// Cantidad promedio de comparaciones para las inserciones realizadas, ponderadas por la altura donde se encuentra el nodo.
	fmt.Printf("Cantidad promedio de comparaciones para las inserciones, ponderada por la altura de cada nodo: %f \n", GetAvgComparisons(ts.tree))
}

//Utiliza el arbol dentro de tree_stats y le inserta todos los elementos del array
func (ts *TreeStats) populate(size int) bool {
	randnums := randgen.GetRandomArray(17, size)
	if len(randnums) >= 200 && len(randnums) <= 1000 {
		for _, value := range randnums {
			ts.insertComparisons += ts.tree.Insert(value)
		}
		ts.comparisons += ts.insertComparisons
		return true
	} else {
		fmt.Println("Verifique que el tamaño del arreglo a insertar está entre 200 y 1000")
		return false
	}
}

// Verifica que el arbol sea de tipo BinaryTree y aplica el balanceo por el algoritmo DSW
func (ts *TreeStats) TransformDSW() {
	switch typedTree := ts.tree.(type) {
	case *BinaryTree:
		typedTree.BalanceDsw()
	default:
	}
}

// Busca todos los números en un slice pasado por parámetro
func (ts *TreeStats) search(slice []int) {
	for _, value := range slice {
		_, comp := ts.tree.Search(value)
		ts.searchComparisons += comp
	}
	ts.comparisons += ts.searchComparisons
}
