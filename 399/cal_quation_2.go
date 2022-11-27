package _99

// 官方题解思想，根据java代码自行开发

type graphNode struct {
	parent []int32
	weight []float64
}

var graph *graphNode

func CalcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	return calcEquation2(equations, values, queries)
}

func calcEquation2(equations [][]string, values []float64, queries [][]string) []float64 {
	graph = &graphNode{
		parent: make([]int32, len(equations)*2),
		weight: make([]float64, len(equations)*2),
	}
	for i := 0; i < len(equations)*2; i++ {
		graph.parent[i] = int32(i)
		graph.weight[i] = 1
	}

	equationMap := map[string]int32{}
	var index int32
	for valueIndex, equation := range equations {
		if _, exist := equationMap[equation[0]]; !exist {
			equationMap[equation[0]] = index
			index++
		}
		if _, exist := equationMap[equation[1]]; !exist {
			equationMap[equation[1]] = index
			index++
		}
		graph.unionNode(equationMap[equation[0]], equationMap[equation[1]], values[valueIndex])
	}

	var res []float64
	for _, query := range queries {
		_, exist1 := equationMap[query[0]]
		_, exist2 := equationMap[query[1]]
		if !exist1 || !exist2 {
			res = append(res, -1)
		} else {
			res = append(res, graph.isConnected(equationMap[query[0]], equationMap[query[1]]))
		}
	}

	return res
}

func (g *graphNode) unionNode(x, y int32, value float64) {
	rootX := g.find(x)
	rootY := g.find(y)

	if rootX == rootY {
		return
	}

	g.parent[rootX] = rootY
	g.weight[rootX] = value * g.weight[y] / g.weight[x]
}

// 边查询边更新
func (g *graphNode) find(x int32) int32 {
	if x != g.parent[x] {
		origin := g.parent[x]
		g.parent[x] = g.find(origin)
		g.weight[x] *= g.weight[origin]
	}

	return graph.parent[x]
}

func (g *graphNode) isConnected(x, y int32) float64 {
	rootX := g.find(x)
	rootY := g.find(y)
	if rootX == rootY {
		return g.weight[x] / g.weight[y]
	}
	return -1.0
}
