package trees

// type Tree struct {
// 	key           int
// 	value         int
// 	leftChildren  *Tree
// 	rightChildren *Tree
// }

type ITree interface {
	Insert(int) int
	Search(int) (bool, int)
	IsEmpty() bool
	IsBranchless() bool
	GetKey() int
	GetValue() int
	GetLeftChildren() ITree
	GetRightChildren() ITree
	IRD()
}

func TreeSearch(tree ITree, toFind int) (bool, int) {
	root := tree
	find := false
	count := 0
	return TreeSearchRec(root, toFind, &find, &count)
}

func TreeSearchRec(tree ITree, toFind int, found *bool, comparations *int) (bool, int) {
	// emptyTree := ITree
	if tree.IsBranchless() {
		*comparations += 1
	} else {
		if tree.GetKey() == toFind {
			*found = true
			*comparations += 2
		} else {
			if toFind <= tree.GetKey() {
				*comparations += 3
				TreeSearchRec(tree.GetLeftChildren(), toFind, found, comparations)
			} else {
				*comparations += 3
				TreeSearchRec(tree.GetRightChildren(), toFind, found, comparations)
			}
		}
	}
	return *found, *comparations
}

