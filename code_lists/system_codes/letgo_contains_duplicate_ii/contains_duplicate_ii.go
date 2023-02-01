package letgo_contains_duplicate_ii

import (
	_ "letgo_repo/system_file/code_enter"
)

/*存在重复元素 II | https://leetcode.cn/problems/contains-duplicate-ii*/

/*给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。*/

func containsNearbyDuplicate(nums []int, k int) bool {
	// k 代表滑动窗口的最大长度
	valueIndexMap := make(map[int]interface{})

	lIndex := 0

	for index, num := range nums {
		if index-lIndex > k {
			delete(valueIndexMap, nums[lIndex])
			lIndex++
		}

		if _, ok := valueIndexMap[num]; ok {
			return ok
		}

		valueIndexMap[num] = struct{}{}
	}

	return false
}
