package letgo_find_the_duplicate_number

import (
	_ "letgo_repo/system_file/code_enter"
)

/*寻找重复数 | https://leetcode.cn/problems/find-the-duplicate-number*/

/*给定一个包含 n + 1 个整数的数组 nums ，其数字都在 [1, n] 范围内（包括 1 和 n），可知至少存在一个重复的整数。

假设 nums 只有 一个重复的整数 ，返回 这个重复的数 。

你设计的解决方案必须 不修改 数组 nums 且只用常量级 O(1) 的额外空间。
*/

func findDuplicate(nums []int) int {
	// 双指针解法,寻找重复的数字
	n := len(nums)
	left, right := 0, n-1

	ans := 0

	for left <= right {
		mid := (right-left)>>1 + left
		count := 0
		for _, num := range nums {
			if num > mid {
				continue
			}
			count++
		}

		if count <= mid {
			left = mid + 1
			continue
		}

		right = mid - 1
		ans = mid
	}

	return ans
}
