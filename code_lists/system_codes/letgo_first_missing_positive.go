package system_codes

/*缺失的第一个正数 | https://leetcode.cn/problems/first-missing-positive*/

// firstMissingPositive
//
//	@Description: 给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
//
// 请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。
//
//	@param nums
//	@return int
func firstMissingPositive(nums []int) int {

	bound := len(nums) + 1

	for index, num := range nums {
		if num <= 0 {
			nums[index] = -nums[index] + bound
		}
	}

	for _, num := range nums {
		if num < 0 {
			num *= -1
		}

		if num >= bound {
			continue
		}

		if nums[num-1] > 0 {
			nums[num-1] *= -1
		}
	}

	for i, num := range nums {
		if num < 0 {
			continue
		}

		return i + 1
	}

	return bound
}
