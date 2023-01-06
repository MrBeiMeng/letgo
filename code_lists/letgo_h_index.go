package code_lists

import "sort"

/*H 指数 | https://leetcode.cn/problems/h-index*/

func hIndex(citations []int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	h := 0
	for i, citation := range citations {
		tmpH := hIndexMin(i+1, citation) // i+1 表示当前统计的论文数量
		h = hIndexMax(h, tmpH)
	}

	return h
}

func hIndexMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func hIndexMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
