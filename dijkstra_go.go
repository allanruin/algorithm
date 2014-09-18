package main

import (
	// "agraph"
	"container/list"
	"fmt"
)

var INF int = 10000

// vnode-arc-arc-arc
// |
// vnode-arc-arc
// |
// vnode

// 边
type Arc struct {
	vertex int // 边指向的终节点
	// nextarc *Arc
	weight int
}

// type Vnode struct {
// 	data     int
// 	firstarc *Arc
// }

type Vnode struct {
	data   int
	arlist *list.List
}

type Agraph struct {
	n       int
	e       int
	adjlist []*Vnode
}

func buildGraph() *Agraph {
	edges := []string{
		"0-1 4",
		"0-2 6",
		"0-3 6",
		"1-4 7",
		"1-2 1",
		"2-4 6",
		"2-5 4",
		"3-2 2",
		"3-5 5",
		"4-6 6",
		"5-6 8",
		"5-4 1",
	}
	graph := Agraph{n: 7, e: len(edges)}
	graph.adjlist = make([]*Vnode, graph.n)

	for _, edge := range edges {
		s := 0
		t := 0
		w := 0

		// 如果要手工输入，只需要稍微更改，并使用Scanf即可
		fmt.Sscanf(edge, "%d-%d %d", &s, &t, &w)
		// fmt.Printf("%d,%d,%d\n", s, t, w)
		if graph.adjlist[s] == nil {
			graph.adjlist[s] = &Vnode{s, list.New()}
		}
		if graph.adjlist[t] == nil {
			graph.adjlist[t] = &Vnode{t, list.New()}
		}

		graph.adjlist[s].arlist.PushBack(&Arc{vertex: t, weight: w})
		// graph.adjlist[s].arlist = append(graph.adjlist[s].arlist, &Arc{vertex: t, weight: w})

	}
	return &graph
}

/*核心算法逻辑*/
func dijkstra(graph *Agraph, dist []int, path []int, s int) {
	dist[s] = 0

	// 字典m用于记录还未被归入已知集合S的点与其（到源点的）距离
	m := make(map[int]int)
	for i, v := range dist {
		m[i] = v
	}

	for x := pick(m); x != -1; x = pick(m) {
		// 对于在未知集合T中距离源点最短的点x，遍历其出发的边，更新可能的距离
		for e := graph.adjlist[x].arlist.Front(); e != nil; e = e.Next() {
			arc := e.Value.(*Arc)

			if dist[x]+arc.weight < dist[arc.vertex] {

				dist[arc.vertex] = dist[x] + arc.weight
				m[arc.vertex] = dist[arc.vertex]
				path[arc.vertex] = x
				// printSlice(dist)
			}
		}
		delete(m, x) // x 归入已知集合S，从未知集合T中删除x的记录
	}
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

func printSlice(sl []int) {
	for _, v := range sl {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}

func printMap(m map[int]int) {
	for k, v := range m {
		fmt.Printf("%d:%d \n", k, v)
	}
	fmt.Printf("\n")
}

func main() {
	fmt.Println("hello")
	// graph := agraph.BuildGraph()
	graph := buildGraph()

	// 源点
	s := 0

	// prepare dist and path
	dist := make([]int, graph.n)
	path := make([]int, graph.n)
	for i, _ := range dist {
		dist[i] = INF
		path[s] = -1
	}

	dijkstra(graph, dist, path, 0)

	fmt.Println("dist[]:")
	printSlice(dist)
	fmt.Println("path[]:")
	printSlice(path)
}
