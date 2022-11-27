package _99

type TreeNode1 struct {
	PointValue string
	EdgeValue  float64
	Children   map[string]*TreeNode1
}

const unknownRes float64 = -1.0

// leetcode 399题

//func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
//	return calcEquation1(equations,values,queries)
//}

// 可以考虑用树来解决这个问题，
func calcEquation1(equations [][]string, values []float64, queries [][]string) []float64 {
	// 把已知条件处理成一颗颗树
	root := &TreeNode1{
		Children: map[string]*TreeNode1{},
	}
	for index, pointValues := range equations {
		pointA, pointB := pointValues[0], pointValues[1]
		value := values[index]
		insertNewNode(root, pointA, pointB, value)
	}

	res := make([]float64, len(queries), len(queries))
	for index, queryComb := range queries {
		res[index] = unknownRes
		queryA := queryComb[0]
		queryB := queryComb[1]
		if find, factors := findChildNode(root, map[string]bool{}, queryA, queryB); find {
			res[index] = calFactor(factors)
		}
	}

	return res
}

func calFactor(factors []float64) float64 {
	res := float64(1)
	for _, factor := range factors {
		res = res * factor
	}
	return res
}

func findChildNode(root *TreeNode1, already map[string]bool, tempPoint, targetPoint string) (bool, []float64) {
	var factors []float64

	if _, exist := root.Children[tempPoint]; !exist {
		return false, nil
	}
	if _, exist := root.Children[targetPoint]; !exist {
		return false, nil
	}

	already[tempPoint] = true
	node := root.Children[tempPoint]
	//如果有的话返回就好
	if temp, exist := node.Children[targetPoint]; exist {
		factors = append(factors, temp.EdgeValue)
		return true, factors
	}
	// 没有的话开始遍历
	for _, childNode := range node.Children {
		if _, exist := already[childNode.PointValue]; exist {
			// 防止无限循环
			continue
		}
		if find, tempFactor := findChildNode(root, already, childNode.PointValue, targetPoint); find {
			factors = append(factors, childNode.EdgeValue)
			factors = append(factors, tempFactor...)
			return find, factors
		}
	}

	return false, nil
}

func insertNewNode(root *TreeNode1, bigChildPoint, littleChildPoint string, childEdge float64) {

	var bigNode *TreeNode1
	var exist bool
	if _, exist = root.Children[bigChildPoint]; !exist {
		root.Children[bigChildPoint] = &TreeNode1{
			PointValue: bigChildPoint,
			EdgeValue:  0,
			Children:   map[string]*TreeNode1{},
		}
	}
	bigNode = root.Children[bigChildPoint]
	bigNode.Children[littleChildPoint] = &TreeNode1{
		PointValue: littleChildPoint,
		EdgeValue:  childEdge,
		Children:   map[string]*TreeNode1{},
	}

	var littleNode *TreeNode1
	exist = false
	if _, exist = root.Children[littleChildPoint]; !exist {
		root.Children[littleChildPoint] = &TreeNode1{
			PointValue: littleChildPoint,
			EdgeValue:  0,
			Children:   map[string]*TreeNode1{},
		}
	}
	littleNode = root.Children[littleChildPoint]
	littleNode.Children[bigChildPoint] = &TreeNode1{
		PointValue: bigChildPoint,
		EdgeValue:  1 / childEdge,
		Children:   map[string]*TreeNode1{},
	}
}
