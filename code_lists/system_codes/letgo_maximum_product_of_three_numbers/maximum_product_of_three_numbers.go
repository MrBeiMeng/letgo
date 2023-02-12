package letgo_maximum_product_of_three_numbers

import (
	_ "letgo_repo/system_file/code_enter"
	"sort"
)

/*三个数的最大乘积 | https://leetcode.cn/problems/maximum-product-of-three-numbers*/

/*给你一个整型数组 nums ，在数组中找出由三个数组成的最大乘积，并输出这个乘积。*/

func maximumProduct(nums []int) int {
	// 警惕符号的坑
	sort.Sort(sort.IntSlice(nums))

	c := nums[len(nums)-1]

	planA := nums[0] * nums[1]

	planB := nums[len(nums)-2] * nums[len(nums)-3]

	return max(planA*c, planB*c)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
