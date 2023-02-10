package letgo_missing_number

import (
	_ "letgo_repo/system_file/code_enter"
)

/*丢失的数字 | https://leetcode.cn/problems/missing-number*/

/*给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。*/

func missingNumber(nums []int) int {

	nums = append(nums, nums[0])
	n := len(nums)
	for _, num := range nums {
		nums[num%n] += n
	}

	for index, num := range nums {
		if num < n {
			return index
		}
	}

	return 0
}
