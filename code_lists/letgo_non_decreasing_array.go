package code_lists

/*非递减数列 | https://leetcode.cn/problems/non-decreasing-array*/

func checkPossibility(nums []int) bool {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			continue
		}
		tmp := nums[i]
		nums[i] = nums[i+1]
		ok1 := checkPossibilityCheckArr(nums)
		nums[i] = tmp
		// 上面重置状态 copy 函数有点不爽
		nums[i+1] = nums[i]
		ok2 := checkPossibilityCheckArr(nums)

		return ok2 || ok1
	}

	return true
}

func checkPossibilityCheckArr(nums []int) bool {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			continue
		}
		return false
	}
	return true
}
