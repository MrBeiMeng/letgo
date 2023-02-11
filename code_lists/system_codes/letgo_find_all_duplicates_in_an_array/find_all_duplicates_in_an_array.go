package letgo_find_all_duplicates_in_an_array

import (
	_ "letgo_repo/system_file/code_enter"
)

/*数组中重复的数据 | https://leetcode.cn/problems/find-all-duplicates-in-an-array*/

/*给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次 的整数，并以数组形式返回。

你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间的算法解决此问题。
*/

func findDuplicates(nums []int) []int {
	bound := len(nums) + 1

	for _, num := range nums {
		nums[num%bound-1] += bound
	}

	answerArr := make([]int, 0)

	for index, num := range nums {
		if num/bound < 2 {
			continue
		}

		answerArr = append(answerArr, index+1)
	}

	return answerArr
}
