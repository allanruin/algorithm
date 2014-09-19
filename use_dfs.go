package main

import (
	"fmt"
	"me.irix/graph"
)

func main() {
	edges := []string{
		"0-1 1",
		"1-2 1",
		"1-3 2",
		"2-4 3",
		"2-5 4",
		"3-6 5",
		"3-7 6",
		"4-8 7",
		"5-8 8",
		"6-8 9",
		"7-8 10",
	}

	mygraph := graph.BuildGraph(edges)

	tmp := make([]bool, mygraph.N)
	visited := &tmp
	seq := graph.DFS_A(mygraph, 0, visited)
	fmt.Printf("seq: %v\n", seq)

	// 构建一个邻接矩阵表示的图，并调用邻接矩阵版本的深度优先算法
	mymgraph := graph.BuildMGraph(edges)
	tmp2 := make([]bool, mymgraph.N)
	visited2 := &tmp2
	seq2 := graph.DFS_M(mymgraph, 0, visited2)
	fmt.Printf("seq: %v\n", seq2)
}
