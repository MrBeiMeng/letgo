package system_codes

/*最大连续 1 的个数 | https://leetcode.cn/problems/max-consecutive-ones*/

func findMaxConsecutiveOnes(nums []int) int {

	i, tooShort := checkLen(nums)
	if tooShort {
		return i
	}

	left, right, answer := searchInMid(nums)

	if left >= 0 { // search in left
		answer = max(findMaxConsecutiveOnes(nums[0:left]), answer)
	}
	if right <= len(nums) { // search in right
		answer = max(findMaxConsecutiveOnes(nums[right:]), answer)
	}

	return answer
}

func searchInMid(nums []int) (int, int, int) {
	answer := 0
	right := len(nums) / 2
	left := right - 1

	for left >= 0 && right < len(nums) {
		if nums[left] == 0 && nums[right] == 0 {
			break
		}

		if nums[left] == 1 {
			left--
			answer++
		}

		if nums[right] == 1 {
			right++
			answer++
		}

	}
	return left, right, answer
}

func checkLen(nums []int) (int, bool) {
	if len(nums) <= 1 {
		if len(nums) == 0 {
			return 0, true
		}

		if nums[0] == 1 {
			return 1, true
		}
		return 0, true
	}
	return 0, false
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

//func findMaxConsecutiveOnes(nums []int) int {
//	biggestSum := 0
//	tmpSum := 0
//
//	for _, num := range nums {
//		if num != 1 {
//			biggestSum = max(tmpSum, biggestSum)
//			tmpSum = 0
//			continue
//		}
//
//		tmpSum += 1
//	}
//
//	biggestSum = max(tmpSum, biggestSum)
//
//	return biggestSum
//}
