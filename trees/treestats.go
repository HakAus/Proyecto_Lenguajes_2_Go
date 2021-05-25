package trees

import (
	"fmt"

	"github.com/Xfighter32/Proyecto_Lenguajes_2_Go/randgen"
)

//Funcion que for( recibe los n y genera los arboles) para cada cantidad y la insercion de todos los elementos
//Se crean 2 arboles binarios ordinarios 1 AVL
//Despues transformar el Binario a DSW
//Generrar una secuencia de 10 000 pseudo aleatorios 1 para los 3 arboles y realizar busquedas

//struct (arbol,comparaciones,altura maxima,altura promedio, promedio de comparaciones ) tree_stats metodos de insercion que recibe
//el arreglo y metodos de medir altura e insertar comparaciones
//arbol.altura

//Metodos de arboles
//Altura de los arboles //Altura de un arbol
//Altura promedio de los arboles // Por definir?
//Promedio comparaciones sum(compariones*nivel)/total

//tree_stats

// *Se generan 15 tree_stat y 5 arreglos*
//tree_stats.start()//Crea los arreglos y los tree_stats con su respectivo arbol por cada n se los inserta a un nuevo tree_stats

//tree_stats.populate_tree(Array)

//*tree_stats.transformDSW()

//tree_stats.searchInTree(array) //Recibe el arreglo de 10 000 e internamente acumulara las comparaciones

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

		// Se genera una secuencia de 10 000 números aleatorios
		biglist := randgen.GetRandomArray(17, 10000)

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
	fmt.Printf("Altura primedio: %f \n", GetAvgHeight(ts.tree))
	// Densidad del árbol
	fmt.Printf("Densidad del árbol: %f \n", Density(ts.tree))
	// Cantidad total de comparaciones realizadas
	fmt.Printf("Cantidad total de comparaciones: %d \n", ts.comparisons)
	// Cantidad promedio de comparaciones para las inserciones realizadas
	fmt.Printf("Cantidad promedio de comparaciones: %f \n", GetAvgComparisons(ts.tree))
}

//Utiliza el arbol dentro de tree_stats y le inserta todos los elementos del array
func (ts TreeStats) populate(size int) {
	randnums := randgen.GetRandomArray(17, size)
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
