package graph

import (
// "container/stack"
// "fmt"
)

// 拓扑排序的图必须要有入度为0的点
func TopologySort(graph *Agraph) []int {
	seq := make([]int, 0)
	// initialize degree array
	degree, zeros := initializeDegreeArray(graph)
	// fmt.Printf("degree:%v,zeros:%v\n", degree, zeros)
	st := Stack{zeros}
	for st.Size() != 0 {
		val := st.Pop()
		seq = append(seq, val)
		// 获取所有相邻的点并把其入度减一
		vs := graph.getAdjs(val)
		// fmt.Printf("getadjs:%v\n", vs)
		for i := range vs {
			t := vs[i]
			degree[t] = degree[t] - 1
			if degree[t] == 0 {
				//有新的点入度为0，扔到入度零的队列zeros中
				st.Push(t)
			}
		}

	}
	return seq
}

func initializeDegreeArray(graph *Agraph) ([]int, []int) {
	degree := make([]int, graph.N)
	zeros := make([]int, 0) // 储存为入度为0的顶点的编号

	for i := range graph.Adjlist {
		for e := graph.Adjlist[i].Arlist.Front(); e != nil; e = e.Next() {
			arc := e.Value.(*Arc)
			degree[arc.Vertex]++
		}
	}

	for i := range degree {
		if degree[i] == 0 {
			zeros = append(zeros, i)
		}
	}

	return degree, zeros
}
