package agraph

import (
	"errors"
)

func Bellmanford(graph *Agraph, s int) ([]int, []int, error) {
	// prepare 函数在dijkstra.go
	dist, path := prepare(graph)

	dist[s] = 0

	for i := 0; i < graph.N-1; i++ {

		// loop all edges and make update
		for j := range graph.Adjlist {
			for e := graph.Adjlist[j].Arlist.Front(); e != nil; e = e.Next() {
				arc := e.Value.(*Arc)
				t := arc.Vertex // 终结点
				if dist[j]+arc.Weight < dist[t] {
					dist[t] = dist[j] + arc.Weight
					path[t] = j
				}
			}
		}

	}

	// check for negative-weight cycles
	// loop all edges and make update
	for j := range graph.Adjlist {
		for e := graph.Adjlist[j].Arlist.Front(); e != nil; e = e.Next() {
			arc := e.Value.(*Arc)
			t := arc.Vertex // 终结点
			if dist[j]+arc.Weight < dist[t] {
				return nil, nil, errors.New("negative-weight cycles found")
			}

		}
	}
	return dist, path, nil
}
