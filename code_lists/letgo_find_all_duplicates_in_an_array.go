package code_lists

/*数组中重复的数据 | https://leetcode.cn/problems/find-all-duplicates-in-an-array*/

// findDuplicates
//
//	@Description:给你一个长度为 n 的整数数组 nums ，其中 nums 的所有整数都在范围 [1, n] 内，且每个整数出现 一次 或 两次 。请你找出所有出现 两次 的整数，并以数组形式返回。
//
// 你必须设计并实现一个时间复杂度为 O(n) 且仅使用常量额外空间的算法解决此问题。
//
//	@param nums
//	@return []int
func findDuplicates(nums []int) []int {

	rightBound := len(nums) + 1
	resultList := make([]int, 0)

	for _, num := range nums {
		num %= rightBound
		nums[num-1] += rightBound

		// 判断
		if nums[num-1]/rightBound >= 2 {
			resultList = append(resultList, num)
		}
	}

	return resultList
}
