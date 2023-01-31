package system_codes

/*错误的集合 | https://leetcode.cn/problems/set-mismatch*/

func findErrorNums(nums []int) []int {

	numValueMap, repeatNum := getValueMap(nums)

	for i := 1; i <= len(nums); i++ {

		if _, ok := numValueMap[i]; !ok {
			return []int{repeatNum, i}
		}

	}

	return nil

}

func getValueMap(nums []int) (map[int]int, int) {
	repeatNum := 0
	numValueMap := make(map[int]int)

	for _, num := range nums {

		if _, ok := numValueMap[num]; ok {
			repeatNum = num
		}

		numValueMap[num] += 1
	}
	return numValueMap, repeatNum
}
