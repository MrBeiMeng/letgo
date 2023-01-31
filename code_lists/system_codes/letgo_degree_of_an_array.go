package system_codes

/*数组的度 | https://leetcode.cn/problems/degree-of-an-array*/

func findShortestSubArray(nums []int) int {

	tmpMap := make(map[int][]int)
	maxTimes := 0

	for index, num := range nums {
		tmpMap[num] = append(tmpMap[num], index)

		maxTimes = findMax(len(tmpMap[num]), maxTimes)
	}

	smallestArrLength := len(nums)
	for _, value := range tmpMap {
		if len(value) < maxTimes {
			continue
		}
		smallestArrLength = findSmall(smallestArrLength, arrLen(value))
	}

	return smallestArrLength
}

func arrLen(value []int) int {
	if len(value) == 1 {
		return 1
	}
	return value[len(value)-1] - value[0] + 1
}

func findSmall(a, b int) int {
	if a <= b {
		return a
	}

	return b
}

func findMax(a, b int) int {
	if a >= b {
		return a
	}

	return b
}
