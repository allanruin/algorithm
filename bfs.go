package graph

import (
	"fmt"
)

func BFS_M(graph *Mgraph, s int) {

}

func BFS_A(graph *Agraph, s int) []int {
	queue := make([]int, 0)
	seq := make([]int, 0)
	visited := make([]bool, graph.N)
	queue = append(queue, s)

	front := 0
	for front <= len(queue)-1 {
		fmt.Printf("front:%v ,queue:%v\n", front, queue)

		seq = append(seq, queue[front])
		visited[front] = true
		// 把正在访问的这个点的相邻点都扔进queue里
		for e := graph.Adjlist[queue[front]].Arlist.Front(); e != nil; e = e.Next() {
			arc := e.Value.(*Arc)
			fmt.Println("arc vertex:", arc.Vertex)
			if visited[arc.Vertex] == false {
				queue = append(queue, arc.Vertex)
			}
		}
		front = front + 1
	}

	return seq
}
