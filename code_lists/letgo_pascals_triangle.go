package code_lists

/*杨辉三角 | https://leetcode.cn/problems/pascals-triangle*/

// generate
//
//	@Description: 给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。
//
// 在「杨辉三角」中，每个数是它左上方和右上方的数的和。
//
//	@param numRows
//	@return [][]int
func generate(numRows int) [][]int {

	baseArr := [][]int{{1}, {1, 1}}

	if numRows <= 2 {
		return baseArr[:numRows]
	}

	for i := 0; i < numRows-2; i++ {

		var tmpArr []int
		tmpArr = append(tmpArr, 1)
		for j := 0; j < len(baseArr[i+1])-1; j++ {
			tmpArr = append(tmpArr, baseArr[i+1][j]+baseArr[i+1][j+1])
		}
		tmpArr = append(tmpArr, 1)
		baseArr = append(baseArr, tmpArr)
	}

	return baseArr
}
