package graph

import (
// "fmt"
)

func MakeHeap(seq *[]int) {
	s := len(*seq) / 2

	for i := s; i > 0; i-- {
		// fmt.Println("adjusting ", i)
		adjust(seq, i, len(*seq))
	}
}

/*基于pickMin的adjust实现*/
func adjust_old(seq *[]int, s int, hi int) {
	lc := s * 2   //left-child
	rc := s*2 + 1 // right-child
	// fmt.Printf("lc:%d,rc:%d,len:%d\n", lc, rc, len(*seq))
	min_index := pickMin((*seq), s, lc, rc, hi)

	// 交换
	if min_index != s {
		tmp := (*seq)[min_index]
		(*seq)[min_index] = (*seq)[s]
		(*seq)[s] = tmp

		// 只有被掉过的那个非s的点，可能会和后代不符合heap要求
		adjust(seq, min_index, hi)
	}
}

/*简化版的实现*/
func adjust(pseq *[]int, s int, hi int) {
	seq := *pseq
	lc := s * 2 //left-child
	p := lc
	// fmt.Printf("lc:%d,len:%d\n", lc, len(seq))

	// 交换
	if lc <= hi-1 {
		// 如果右孩子存在，且比左孩子小，则p指向右孩子。既P是指向孩子中较小的那个
		if lc+1 <= hi-1 && seq[lc] > seq[lc+1] {
			p = lc + 1
		}

		// 如果孩子比点s小
		if seq[p] < seq[s] {
			tmp := seq[p]
			seq[p] = seq[s]
			seq[s] = tmp

			// 只有被掉过的那个非s的点，可能会和后代不符合heap要求
			adjust(pseq, p, hi)
		}

	}

}

// return the index of the minimum element
// 关键问题是由于通过2*s 来求得的l,r，所以有可能超过了slice的长度
// 要做判断根据是否有右孩子、是否有左孩子进行不同的判断
// 通过将挑选函数单独出来，可以方便换成pickMax就可以转成大项堆了
func pickMin(ar []int, s, l, r, hi int) int {
	// 最后一个下标
	// le := len(ar) - 1
	le := hi - 1

	if r > le && l > le {
		return s
	}

	// 太多边界条件容易出错了，比如这里如果是l<le就错了
	if r > le && l <= le {
		if ar[l] < ar[s] {
			return l
		} else {
			return s
		}
	}

	if ar[l] < ar[s] {
		if ar[l] < ar[r] {
			return l
		} else {
			return r
		}
	} else {
		if ar[s] < ar[r] {
			return s
		} else {
			return r
		}
	}
}

func pickMin2(ar []int, l int, r int) int {
	if ar[l] < ar[r] {
		return l
	} else {
		return r
	}
}

// 排序成为大->小
func Heapsort(pseq *[]int) {
	MakeHeap(pseq)
	seq := *pseq

	for i := 1; i <= len(seq)-1; i++ {
		tmp := seq[1]
		seq[1] = seq[len(seq)-i]
		seq[len(seq)-i] = tmp
		adjust(pseq, 1, len(seq)-i)
		// fmt.Printf("%v\n", seq)
	}

}
