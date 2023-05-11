package letgo_3sum

import (
	_ "letgo_repo/system_file/code_enter"
	"sort"
)

/*三数之和 | https://leetcode.cn/problems/3sum*/

/*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
示例 3：

输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。


提示：

3 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

func threeSum(nums []int) [][]int {
	// 将三数之和变为两数之和
	answer := make([][]int, 0)

	// 排序
	sort.Ints(nums)

	// 固定一个元素，查找两数之和
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]

		left, right := i+1, len(nums)-1

		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				answer = append(answer, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum < target {
				left++
			} else {
				right--
			}

		}

	}

	return answer
}

// twoSum
//
//	@Description: 查找所有的两数之和为target 的可能
//	@param nums
//	@param target
//	@return []int
//	@return bool
//func twoSum(nums []int, target int) ([][]int, bool) {
//	valueMap := make(map[int]struct{})
//	answer := make([][]int, 0)
//
//	for _, num1 := range nums {
//		num2 := target - num1
//		if _, ok := valueMap[num2]; ok {
//
//			if num1 <= num2 {
//				answer = append(answer, []int{num1, num2})
//			} else {
//				answer = append(answer, []int{num2, num1})
//			}
//
//		}
//
//		valueMap[num1] = struct{}{}
//	}
//
//	return answer, len(answer) > 0
//}
