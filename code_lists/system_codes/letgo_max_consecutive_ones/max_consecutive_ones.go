package letgo_max_consecutive_ones

import (
	_ "letgo_repo/system_file/code_enter"
)

/*最大连续 1 的个数 | https://leetcode.cn/problems/max-consecutive-ones*/

/*给定一个二进制数组 nums ， 计算其中最大连续 1 的个数。*/

func findMaxConsecutiveOnes(nums []int) int {
	answer, dp := 0, 0

	for i, num := range nums {
		if num == 1 {
			dp++

			if i < len(nums)-1 { // 最后一个是1也要判断一下
				continue
			}
		}

		answer = max(answer, dp)
		dp = 0
	}

	return answer
}
func max(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
