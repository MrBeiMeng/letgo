package system_codes

import "sort"

/*最小操作次数使数组元素相等 | https://leetcode.cn/problems/minimum-moves-to-equal-array-elements*/

// minMoves
//
//	@Description: 给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。
//	@param nums
//	@return int
func minMoves(nums []int) (result int) {

	sort.Sort(sort.IntSlice(nums))

	n := len(nums) - 1 //去除第一位

	for i, num := range nums {
		if i == 0 {
			continue
		}

		result += n * (num - nums[i-1])
		n--
	}
	return
}
