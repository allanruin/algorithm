package agraph

import (
	"fmt"
	"reflect"
	"testing"
)

// negative case
var negativeGraphTests = [][]string{
	// case in wikipedia page about bellman-ford algorithm
	// http://en.wikipedia.org/wiki/Bellman%E2%80%93Ford_algorithm
	[]string{
		"0-1 -3",
		"1-2 1",
		"2-3 1",
		"3-4 1",
	},

	// http://www.geeksforgeeks.org/dynamic-programming-set-23-bellman-ford-algorithm/
	[]string{
		"0-1 -1",
		"0-2 4",
		"1-2 3",
		"1-4 2",
		"1-3 2",
		"3-1 1",
		"4-3 -3",
		"3-2 5",
	},
}

var negativeCycleCases = [][]string{
	// http://cs.stackexchange.com/questions/6919/getting-negative-cycle-using-bellman-ford
	[]string{
		"0-1 -20",
		"1-2 10",
		"2-0 5",
		"2-1 -15",
	},
}

var NegativeagraphResults = []aresult{
	aresult{
		dist: []int{0, -3, -2, -1, 0},
		path: []int{-1, 0, 1, 2, 3},
	},
	aresult{
		dist: []int{0, -1, 2, -2, 1},
		path: []int{-1, 0, 1, 4, 1},
	},
}

func TestBellmanFord(t *testing.T) {
	for i, tt := range agraphTests {
		graph := BuildGraph(tt)
		dist, path, _ := Bellmanford(graph, 0)
		result := aresult{dist: dist, path: path}

		eq := reflect.DeepEqual(agraphResults[i], result)
		if !eq {
			t.Fatalf("第%d个测试失败", i+1)

		}
	}

	for i, tt := range negativeGraphTests {
		graph := BuildGraph(tt)
		dist, path, _ := Bellmanford(graph, 0)
		result := aresult{dist: dist, path: path}

		eq := reflect.DeepEqual(NegativeagraphResults[i], result)
		if !eq {
			fmt.Printf("%v\n", dist)
			fmt.Printf("%v\n", path)
			t.Fatalf("第%d个负边测试失败", i+1)

		}
	}

	// test recognition of negative cycle
	for i, tt := range negativeCycleCases {
		graph := BuildGraph(tt)
		_, _, ok := Bellmanford(graph, 0)
		if ok == nil {
			t.Fatalf("第%d个测试没有识别出负边循环", i+1)
		}

	}

}
