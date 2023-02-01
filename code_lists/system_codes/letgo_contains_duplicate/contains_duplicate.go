package letgo_contains_duplicate

import (
	_ "letgo_repo/system_file/code_enter"
	"sort"
)

/*存在重复元素 | https://leetcode.cn/problems/contains-duplicate*/

/*给你一个整数数组 nums 。如果任一值在数组中出现 至少两次 ，返回 true ；如果数组中每个元素互不相同，返回 false 。*/

func containsDuplicate(nums []int) bool {
	// 两种解法，一种使用哈希表，一种则是进行排序

	sort.Sort(sort.IntSlice(nums))

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}

	}

	return false
}
