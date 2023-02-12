package letgo_degree_of_an_array

import (
	_ "letgo_repo/system_file/code_enter"
)

/*数组的度 | https://leetcode.cn/problems/degree-of-an-array*/

func findShortestSubArray(nums []int) int {
	intervalMap := make(map[int][]int)

	for i, num := range nums {
		if _, ok := intervalMap[num]; ok {
			intervalMap[num] = append(intervalMap[num], i)
			continue
		}
		intervalMap[num] = []int{i}
	}

	ans := len(nums)
	maxD := 0
	for _, value := range intervalMap {
		maxD = max(maxD, len(value))
	}

	if maxD == 1 {
		return 1
	}

	for _, value := range intervalMap {

		if len(value) == maxD {
			ans = min(ans, value[len(value)-1]-value[0]+1)
		}
	}

	return ans
}

func min(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
