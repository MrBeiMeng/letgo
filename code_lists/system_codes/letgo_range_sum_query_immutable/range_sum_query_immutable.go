package letgo_range_sum_query_immutable

import (
	_ "letgo_repo/system_file/code_enter"
)

/*区域和检索 - 数组不可变 | https://leetcode.cn/problems/range-sum-query-immutable*/

type NumArray struct {
	Nums []int
}

func Constructor(nums []int) NumArray {
	var na NumArray
	na.Nums = append(na.Nums, nums...)
	return na
}

func (this *NumArray) SumRange(left int, right int) int {

	sum := 0
	for i := left; i <= right; i++ {
		sum += this.Nums[i]
	}

	return sum
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
