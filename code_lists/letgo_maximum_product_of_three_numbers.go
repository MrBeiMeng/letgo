package code_lists

import "sort"

/*三个数的最大乘积 | https://leetcode.cn/problems/maximum-product-of-three-numbers*/

// maximumProduct
//
//	@Description: 给你一个整型数组 nums ，在数组中找出由三个数组成的最大乘积，并输出这个乘积。
//
// 3 <= nums.length <= 104
// -1000 <= nums[i] <= 1000
// @param nums
// @return int
func maximumProduct(nums []int) int {

	sort.Sort(sort.Reverse(sort.IntSlice(nums)))

	return max(nums[0]*nums[1]*nums[2], nums[0]*nums[len(nums)-2]*nums[len(nums)-1])
}
