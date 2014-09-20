package graph

func MakeHeap(seq *[]int) {
	s := len(*seq) / 2

	for i := s; i > 0; i-- {
		adjust(seq, i)
	}
}

func adjust(seq *[]int, s int) {
	lc := s * 2   //left-child
	rc := s*2 + 1 // right-child
	min_index := pickMin((*seq), s, lc, rc)

	// 交换
	if min_index != s {
		tmp := (*seq)[min_index]
		(*seq)[min_index] = (*seq)[s]
		(*seq)[s] = tmp

		// 只有被掉过的那个非s的点，可能会和后代不符合heap要求
		adjust(seq, min_index)
	}
}

// return the index of the minimum element
func pickMin(ar []int, s, l, r int) int {
	le := len(ar) - 1

	if r > le && l > le {
		return s
	}

	if r > le && le > l {
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
