package graph

import (
	"container/list"
	"fmt"
)

var INF int = 1000000

type Arc struct {
	Vertex int // 边指向的终节点
	Weight int
}

type Vnode struct {
	Data   int
	Arlist *list.List
}

type Agraph struct {
	N       int
	E       int
	Adjlist []*Vnode
}

/*
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
*/
func BuildGraph(edges []string) *Agraph {

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
	graph := Agraph{N: n, E: len(edges)}
	graph.Adjlist = make([]*Vnode, graph.N)

	s = 0
	t = 0
	w = 0
	for _, edge := range edges {

		// 如果要手工输入，只需要稍微更改，并使用Scanf即可
		fmt.Sscanf(edge, "%d-%d %d", &s, &t, &w)
		// fmt.Printf("%d,%d,%d\n", s, t, w)
		if graph.Adjlist[s] == nil {
			graph.Adjlist[s] = &Vnode{s, list.New()}
		}
		if graph.Adjlist[t] == nil {
			// defer fmt.Println("panic t", t)
			graph.Adjlist[t] = &Vnode{t, list.New()}
		}

		graph.Adjlist[s].Arlist.PushBack(&Arc{Vertex: t, Weight: w})

	}
	return &graph
}

func (graph *Agraph) getAdjs(s int) []int {
	seq := make([]int, 0)
	for e := graph.Adjlist[s].Arlist.Front(); e != nil; e = e.Next() {
		arc := e.Value.(*Arc)
		seq = append(seq, arc.Vertex)
	}
	return seq
}

func PrintSlice(sl []int) {
	for _, v := range sl {
		fmt.Printf("%d ", v)
	}
	fmt.Printf("\n")
}

func PrintMap(m map[int]int) {
	for k, v := range m {
		fmt.Printf("%d:%d \n", k, v)
	}
	fmt.Printf("\n")
}
