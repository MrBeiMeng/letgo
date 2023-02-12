package letgo_product_of_array_except_self

import (
	_ "letgo_repo/system_file/code_enter"
)

/*除自身以外数组的乘积 | https://leetcode.cn/problems/product-of-array-except-self*/

// 思考前缀和问题

func productExceptSelf(nums []int) []int {
	tmpChick := 1
	otherChickArr := make([]int, 0)

	for _, num := range nums {
		otherChickArr = append(otherChickArr, tmpChick)
		tmpChick *= num
	}

	tmpChick = 1
	for i := len(nums) - 1; i >= 0; i-- {
		otherChickArr[i] *= tmpChick
		tmpChick *= nums[i]
	}

	return otherChickArr
}
