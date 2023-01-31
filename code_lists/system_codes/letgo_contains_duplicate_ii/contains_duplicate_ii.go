package letgo_contains_duplicate_ii

import (
	_ "letgo_repo/system_file/code_enter"
)

/*存在重复元素 II | https://leetcode.cn/problems/contains-duplicate-ii*/

/*给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。*/

func containsNearbyDuplicate(nums []int, k int) bool {
	valueIndexMap := make(map[int]int)

	minInterval := -1

	for i, value := range nums {
		if index, ok := valueIndexMap[value]; ok {
			minInterval = min(minInterval, abs(i-index))
		}

		valueIndexMap[value] = i
	}

	return minInterval != -1 && minInterval <= k
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}

func min(a, b int) int {
	if a == -1 {
		return b
	}

	if b == -1 {
		return a
	}

	if a <= b {
		return a
	}

	return b
}
