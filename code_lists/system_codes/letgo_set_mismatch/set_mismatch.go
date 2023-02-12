package letgo_set_mismatch

import (
	_ "letgo_repo/system_file/code_enter"
)

/*错误的集合 | https://leetcode.cn/problems/set-mismatch*/

/*集合 s 包含从 1 到n的整数。不幸的是，因为数据错误，导致集合里面某一个数字复制了成了集合里面的另外一个数字的值，导致集合 丢失了一个数字 并且 有一个数字重复 。

给定一个数组 nums 代表了集合 S 发生错误后的结果。

请你找出重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。
*/

func findErrorNums(nums []int) []int {
	bound := len(nums) + 1
	for _, num := range nums {
		nums[num%bound-1] += bound
	}

	a, b := 0, 0

	for index, num := range nums {
		if num < bound {
			b = index + 1
		}

		if num/bound >= 2 {
			a = index + 1
		}
	}

	return []int{a, b}
}
