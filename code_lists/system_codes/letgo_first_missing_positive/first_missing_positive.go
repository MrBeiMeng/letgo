package letgo_first_missing_positive

import (
	_ "letgo_repo/system_file/code_enter"
)

/*缺失的第一个正数 | https://leetcode.cn/problems/first-missing-positive*/

/*给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。

请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。*/

func firstMissingPositive(nums []int) int {
	n := len(nums) + 1
	for index, num := range nums {
		if num < 0 || num > len(nums) {
			nums[index] = 0
		}
	}

	for _, num := range nums {
		// 不和规则的跳过
		if num%n == 0 {
			continue
		}

		tmpIndex := (num % n) - 1
		nums[tmpIndex] += n
	}

	for i, num := range nums {
		if num < n {
			return i + 1
		}
	}

	return n
}
