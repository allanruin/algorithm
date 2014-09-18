package agraph

import (
	"reflect"
	"testing"
)

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
}
