package graph

import (
	"reflect"
	"testing"
)

var stackTests = [][]int{
	[]int{1},
	[]int{1, 2},
	[]int{1, 2, 3},
	[]int{1, 2, 3, 4},
}

var pushResults = [][]int{
	[]int{1, 5},
	[]int{1, 2, 5},
	[]int{1, 2, 3, 5},
	[]int{1, 2, 3, 4, 5},
}

var popResults = [][]int{
	[]int{},
	[]int{1},
	[]int{1, 2},
	[]int{1, 2, 3},
}

func TestStack(t *testing.T) {
	for i, tt := range stackTests {

		st1 := &Stack{tt}
		st1.Push(5)
		eq1 := reflect.DeepEqual(pushResults[i], st1.GetValues())
		if !eq1 {
			t.Errorf("第%d个push测试失败,期待:%v,得到:%v", i+1, pushResults[i], st1.GetValues())
		}

		st2 := &Stack{tt}
		st2.Pop()
		eq2 := reflect.DeepEqual(popResults[i], st2.GetValues())
		if !eq2 {
			t.Errorf("第%d个pop测试失败,期待:%v,得到:%v", i+1, popResults[i], st2.GetValues())
		}

	}
}
