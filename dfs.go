package graph

import (
// "fmt"
)

// 邻接矩阵图的深度遍历
func DFS_M(graph *Mgraph, s int, visited *[]bool) []int {
	seq := make([]int, 0)

	if (*visited)[s] != true {
		(*visited)[s] = true

		seq = append(seq, s)

		// 对每个节点的射出的边，都进行深度优先遍历
		for j := 0; j < graph.N; j++ {
			if graph.Edges[s][j] != INF {
				nseq := DFS_M(graph, j, visited)
				seq = append(seq, nseq...)
			}
		}

	}

	return seq

}

// 邻接表图的深度遍历
// 为了而在每次迭代时共享变量vistied，用以判断哪些已经访问了
func DFS_A(graph *Agraph, s int, visited *[]bool) []int {
	//visited := &make([]bool, graph.N)
	seq := make([]int, 0)

	if (*visited)[s] != true {
		(*visited)[s] = true
		// fmt.Println("visiting ", i)

		seq = append(seq, graph.Adjlist[s].Data)

		// 对每个节点的射出的边，都进行深度优先遍历
		for e := graph.Adjlist[s].Arlist.Front(); e != nil; e = e.Next() {
			arc := e.Value.(*Arc)
			nseq := DFS_A(graph, arc.Vertex, visited) // next sequence
			seq = append(seq, nseq...)
		}

	}

	return seq
}
