package graph

import (
	"fmt"
)

type Mgraph struct {
	Edges [][]int
	N     int
	E     int
}

func BuildMGraph(edges []string) *Mgraph {

	// 自动计算点的个数
	s := 0
	t := 0
	w := 0
	// 没有set容器只好拿map来凑了
	vertexm := make(map[int]bool)
	for _, edge := range edges {

		fmt.Sscanf(edge, "%d-%d %d", &s, &t, &w)
		vertexm[s] = true
		vertexm[t] = true
	}
	n := len(vertexm)

	graph := Mgraph{N: n, E: len(edges)}

	// allocate composed 2d array
	graph.Edges = make([][]int, n)
	for i := range graph.Edges {
		graph.Edges[i] = make([]int, n)
	}

	// intialize to INF
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				graph.Edges[i][j] = INF
			} else {
				graph.Edges[i][j] = 0
			}

		}
	}

	s = 0
	t = 0
	w = 0
	for _, edge := range edges {

		// 如果要手工输入，只需要稍微更改，并使用Scanf即可
		fmt.Sscanf(edge, "%d-%d %d", &s, &t, &w)
		// fmt.Printf("%d,%d,%d\n", s, t, w)
		graph.Edges[s][t] = w

	}
	return &graph
}

func makeArray2d(row int, col int) *[][]int {
	// allocate composed 2d array
	a := make([][]int, row)
	e := make([]int, row*col)
	for i := range a {
		a[i] = e[i*col : (i+1)*col]
	}

	return &a
}

func initializeArray2d(array *[][]int, val int) *[][]int {
	n := len((*array))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			(*array)[i][j] = val
		}
	}
	return array
}

func initializeArray2d_diag(array *[][]int, val int) *[][]int {
	n := len((*array))
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i != j {
				(*array)[i][j] = INF
			} else {
				(*array)[i][j] = 0
			}

		}
	}
	return array
}
