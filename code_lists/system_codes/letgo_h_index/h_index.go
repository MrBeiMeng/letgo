package letgo_h_index

import (
	_ "letgo_repo/system_file/code_enter"
	"sort"
)

/*H 指数 | https://leetcode.cn/problems/h-index*/

func hIndex(citations []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	h := 0

	for count, num := range citations {
		tmpH := min(num, count+1)

		h = max(h, tmpH)
	}

	return h
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
