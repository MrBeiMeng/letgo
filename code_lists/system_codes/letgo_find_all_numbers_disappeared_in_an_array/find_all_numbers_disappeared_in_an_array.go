package letgo_find_all_numbers_disappeared_in_an_array

import (
	_ "letgo_repo/system_file/code_enter"
)

/*找到所有数组中消失的数字 | https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array*/

/*给你一个含 n 个整数的数组 nums ，其中 nums[i] 在区间 [1, n] 内。请你找出所有在 [1, n] 范围内但没有出现在 nums 中的数字，并以数组的形式返回结果。

示例 1：

输入：nums = [4,3,2,7,8,2,3,1]
输出：[5,6]
示例 2：

输入：nums = [1,1]
输出：[2]
*/

func findDisappearedNumbers(nums []int) []int {
	for _, num := range nums {
		if num < 0 {
			num *= -1
		}

		if nums[num-1] < 0 {
			continue
		}

		nums[num-1] *= -1
	}

	ansArr := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			continue
		}

		ansArr = append(ansArr, i+1)
	}

	return ansArr
}
