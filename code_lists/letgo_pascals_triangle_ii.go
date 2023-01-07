package code_lists

/*杨辉三角 II | https://leetcode.cn/problems/pascals-triangle-ii*/

// getRow
//
//	@Description:给定一个非负索引 rowIndex，返回「杨辉三角」的第 rowIndex 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//
//	@param rowIndex
//	@return []int
func getRow(rowIndex int) []int {

	baseArr := [][]int{{1}}

	for i := 0; i < rowIndex; i++ {

		tmpArr := make([]int, 0)
		tmpArr = append(tmpArr, 1)

		lastArr := baseArr[len(baseArr)-1]
		for j := 0; j < len(lastArr)-1; j++ {
			tmpArr = append(tmpArr, lastArr[j]+lastArr[j+1])
		}

		tmpArr = append(tmpArr, 1)

		baseArr = append(baseArr, tmpArr)
	}

	return baseArr[len(baseArr)-1]
}
