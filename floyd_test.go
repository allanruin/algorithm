package graph

import (
	"fmt"
	"reflect"
	"testing"
)

var floydTests = [][]string{
	[]string{
		"0-1 5",
		"0-3 7",
		"1-3 2",
		"1-2 4",
		"3-2 1",
		"2-0 3",
		"2-3 2",
		"2-1 3",
	},
}

type Floydresult struct {
	dist [][]int
	path [][]int
}

var floydResults = []Floydresult{
	Floydresult{
		dist: [][]int{
			[]int{0, 5, 8, 7},
			[]int{6, 0, 3, 2},
			[]int{3, 3, 0, 2},
			[]int{4, 4, 1, 0},
		},
		path: [][]int{
			[]int{-1, -1, 3, -1},
			[]int{3, -1, 3, -1},
			[]int{-1, -1, -1, -1},
			[]int{2, 2, -1, -1},
		},
	},
}

// 测试矩阵表示的dijkstra算法
func TestFloyd(t *testing.T) {
	for i, tt := range floydTests {
		graph := BuildMGraph(tt)
		dist, path := Floyd(graph)
		result := Floydresult{dist: *dist, path: *path}

		eq := reflect.DeepEqual(floydResults[i], result)
		if !eq {
			t.Errorf("第%d个测试失败", i+1)
			fmt.Printf("%v\n", dist)
			fmt.Printf("%v\n", path)
		}
	}
}
