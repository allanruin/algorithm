package graph

import (
	// "reflect"
	"testing"
)

var pickMinTests = [][]int{
	[]int{1, 2, 3},
	[]int{1, 3, 2},
	[]int{3, 1, 2},
	[]int{3, 2, 1},
	[]int{2, 1, 3},
	[]int{2, 3, 1},
}

var pickMinResults = []int{1, 1, 2, 3, 2, 3}

func TestPickMin(t *testing.T) {
	for i, tt := range pickMinTests {

		re := pickMin(tt, 0, 1, 2) + 1
		if re != pickMinResults[i] {
			t.Errorf("第%d个测试失败,期待:%v,得到:%v", i+1, pickMinResults[i], re)
		}

	}
}

// var heapTests = [][]int{
// 	[]int{INF, 49, 38, 65, 97, 76, 13, 27, 49},
// 	[]int{INF, 2, 7, 10, 8, 9, 1},
// }

// var heapResults = [][]int{
// 	[]int{13, 38, 27, 49, 76, 65, 49, 97},
// 	[]int{1, 7, 2, 8, 9, 10},
// }

// func TestMakeHeap(t *testing.T) {
// 	for i, tt := range pickMinTests {

// 		MakeHeap(&tt)
// 		eq := reflect.DeepEqual(tt, heapResults[i])
// 		if !eq {
// 			t.Errorf("第%d个测试失败,期待:%v,得到:%v", i+1, heapResults[i], tt)
// 		}

// 	}
// }
