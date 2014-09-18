package agraph

import (
// "fmt"
)

func prepare(graph *Agraph) ([]int, []int) {
	// prepare dist and path
	dist := make([]int, graph.N)
	path := make([]int, graph.N)
	for i, _ := range dist {
		dist[i] = INF
		path[i] = -1
	}
	return dist, path
}

/*核心算法逻辑*/
func Dijkstra(graph *Agraph, s int) ([]int, []int) {
	dist, path := prepare(graph)

	dist[s] = 0

	// 字典m用于记录还未被归入已知集合S的点与其（到源点的）距离
	m := make(map[int]int)
	for i, v := range dist {
		m[i] = v
	}

	for x := pick(m); x != -1; x = pick(m) {
		// 对于在未知集合T中距离源点最短的点x，遍历其出发的边，更新可能的距离
		for e := graph.Adjlist[x].Arlist.Front(); e != nil; e = e.Next() {
			arc := e.Value.(*Arc)

			if dist[x]+arc.Weight < dist[arc.Vertex] {

				dist[arc.Vertex] = dist[x] + arc.Weight
				m[arc.Vertex] = dist[arc.Vertex]
				path[arc.Vertex] = x
				// printSlice(dist)
			}
		}
		delete(m, x) // x 归入已知集合S，从未知集合T中删除x的记录
	}

	return dist, path
}

// 挑出字典中值最小的key
func pick(m map[int]int) int {
	min := INF
	key := -1
	for k, v := range m {
		if v < min {
			min = v
			key = k
		}
	}
	// fmt.Println("pick ", key)
	// printMap(m)
	return key
}
