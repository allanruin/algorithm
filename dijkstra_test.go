package graph

import (
	"reflect"
	"testing"
)

var agraphTests = [][]string{
	[]string{
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
	},
	[]string{
		// http://www.coder4.com/archives/3506
		"0-2 10",
		"0-4 30",
		"0-5 100",
		"1-2 5",
		"2-3 50",
		"3-5 10",
		"4-3 20",
		"4-5 60",
	},
	[]string{
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
	},
}

type aresult struct {
	dist []int
	path []int
}

var agraphResults = []aresult{
	aresult{
		dist: []int{0, 4, 5, 6, 10, 9, 16},
		path: []int{-1, 0, 1, 0, 5, 2, 4},
	},
	aresult{
		dist: []int{0, INF, 10, 50, 30, 60},
		path: []int{-1, -1, 0, 4, 0, 3},
	},
	aresult{
		dist: []int{0, 20, 5, 10, 56, 97, 61, 84, 9, 73, 78},
		path: []int{-1, 0, 0, 2, 1, 4, 1, 10, 0, 6, 6},
	},
}

func TestDijkstra(t *testing.T) {
	for i, tt := range agraphTests {
		graph := BuildGraph(tt)
		dist, path := Dijkstra(graph, 0)
		result := aresult{dist: dist, path: path}

		eq := reflect.DeepEqual(agraphResults[i], result)
		if !eq {
			t.Errorf("第%d个测试失败", i+1)

		}
	}
}

// 测试邻接图表示的dijkstra算法
func TestDijkstra_M(t *testing.T) {
	for i, tt := range agraphTests {
		graph := BuildMGraph(tt)
		dist, path := Dijkstra_M(graph, 0)
		result := aresult{dist: dist, path: path}

		eq := reflect.DeepEqual(agraphResults[i], result)
		if !eq {
			t.Errorf("第%d个测试失败", i+1)

		}
	}
}
