package main

import (
	"fmt"
	"me.irix/graph"
)

func main() {
	edges := []string{
		"0-1 1",
		"0-2 1",
		"0-3 1",
		"1-2 1",
		"1-4 1",
		"2-4 1",
		"2-5 1",
		"3-5 1",
		"5-4 1",
		"4-6 1",
		"5-6 1",
	}

	mygraph := graph.BuildGraph(edges)
	seq := graph.TopologySort(mygraph)
	fmt.Println("seq:", seq)
	// seq: [0 3 1 2 5 4 6]
}
