package code_lists

import "sort"

/*缺失的第一个正数 | https://leetcode.cn/problems/first-missing-positive*/

// firstMissingPositive
//
//	@Description: 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//
//	@param nums
//	@return int
func firstMissingPositive(nums []int) int {
	sort.Sort(sort.IntSlice(nums))

	big := 0
	for _, num := range nums {
		if num-big == 1 {
			big = num
		}
	}

	return big + 1
}
