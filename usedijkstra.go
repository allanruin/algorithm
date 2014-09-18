package main

import (
	"agraph"
	"fmt"
)

func main() {
	edges := []string{
		"0-8 9",
		"0-1 20",
		"0-2 5",
		"1-2 12",
		"1-4 36",
		"1-8 25",
		"1-6 41",
		"2-3 5",
		"2-4 56",
		"3-9 68",
		"4-5 41",
		"4-9 68",
		"5-10 21",
		"6-9 12",
		"6-10 17",
		"7-6 65",
		"7-5 82",
		"8-2 31",
		"9-8 8",
		"10-3 12",
		"10-7 6",
	}
	graph := agraph.BuildGraph(edges)
	dist, path := agraph.Dijkstra(graph, 0)
	fmt.Printf("dist: %v\n", dist)
	fmt.Printf("path: %v\n", path)
}
