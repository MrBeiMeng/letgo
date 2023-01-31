package system_codes

/*移动零 | https://leetcode.cn/problems/move-zeroes*/

// moveZeroes
//
//	@Description:给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 请注意 ，必须在不复制数组的情况下原地对数组进行操作。
//
//	@param nums
func moveZeroes(nums []int) {

	zeroNum := 0

	for i, num := range nums {
		if num == 0 {
			zeroNum += 1
			continue
		}

		if i-zeroNum >= 0 {
			nums[i-zeroNum] = nums[i]
		}
	}

	for i := 0; i < zeroNum; i++ {
		nums[len(nums)-1-i] = 0
	}
}
