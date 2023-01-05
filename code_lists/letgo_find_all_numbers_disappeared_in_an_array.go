package code_lists

/*找到所有数组中消失的数字 | https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array*/

func findDisappearedNumbers(nums []int) []int {

	resultList := make([]int, 0)
	tmpMap := make(map[int]int)

	for _, num := range nums {
		tmpMap[num] += 1
	}

	for i := 1; i <= len(nums); i++ {
		if _, ok := tmpMap[i]; !ok {
			resultList = append(resultList, i)
		}
	}

	return resultList

}

func max448(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
