package graph

import (
// "fmt"
)

func BFS_M(graph *Mgraph, s int) []int {
	queue := make([]int, 0)
	seq := make([]int, 0)
	set := make([]bool, graph.E)

	queue = append(queue, s)
	set[s] = true

	front := 0
	for front <= len(queue)-1 {
		// fmt.Printf("front:%v ,queue:%v\n", front, queue)

		seq = append(seq, queue[front])

		// 把正在访问的这个点的相邻点都扔进queue里
		t := queue[front]
		for i := 0; i < graph.N; i++ {
			if graph.Edges[t][i] != INF && set[i] == false {
				queue = append(queue, i)
				set[i] = true
			}
		}

		front = front + 1
	}

	return seq
}

func BFS_A(graph *Agraph, s int) []int {
	queue := make([]int, 0)
	seq := make([]int, 0)
	set := make([]bool, graph.E)

	queue = append(queue, s)
	set[s] = true

	front := 0
	for front <= len(queue)-1 {
		// fmt.Printf("front:%v ,queue:%v\n", front, queue)

		seq = append(seq, queue[front])

		// 把正在访问的这个点的相邻点都扔进queue里
		for e := graph.Adjlist[queue[front]].Arlist.Front(); e != nil; e = e.Next() {
			arc := e.Value.(*Arc)
			// fmt.Println("arc vertex:", arc.Vertex)
			if set[arc.Vertex] == false {
				queue = append(queue, arc.Vertex)
				set[arc.Vertex] = true
			}
		}
		front = front + 1
	}

	return seq
}
