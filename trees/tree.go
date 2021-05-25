package trees

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
	if tree.IsEmpty() {
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

func Height(tree ITree) int {
	if tree.IsEmpty() {
		return 0
	} else {
		return 1 + max(Height(tree.GetLeftChildren()), Height(tree.GetRightChildren()))
	}
}

func max(left int, right int) int {
	if left >= right {
		return left
	} else {
		return right
	}
}

func GetAvgHeight(tree ITree) float64 {
	sum := 0
	GetAvgHeight_Rec(tree, 0, &sum)
	return float64(sum / CountNodes(tree))
}

func GetAvgHeight_Rec(tree ITree, level int, sum *int) {
	if !tree.IsEmpty() {
		level++
		*sum += level
		GetAvgHeight_Rec(tree.GetLeftChildren(), level, sum)
		GetAvgHeight_Rec(tree.GetRightChildren(), level, sum)
	}
}

func GetAvgComparisons(tree ITree) float64 {
	sum := 0
	GetAvgComparisons_Rec(tree, &sum, 0)
	return float64(sum) / float64(CountNodes(tree))
}

func GetAvgComparisons_Rec(tree ITree, sum *int, level int) {
	if !tree.IsEmpty() {
		level++
		*sum += level * tree.GetValue()
		GetAvgComparisons_Rec(tree.GetLeftChildren(), sum, level)
		GetAvgComparisons_Rec(tree.GetRightChildren(), sum, level)
	}
}

func CountNodes(tree ITree) int {
	if tree.IsEmpty() {
		return 0
	}
	return 1 + CountNodes(tree.GetLeftChildren()) + CountNodes(tree.GetRightChildren())
}

func Density(tree ITree) float64 {
	return float64(CountNodes(tree) / Height(tree))
}