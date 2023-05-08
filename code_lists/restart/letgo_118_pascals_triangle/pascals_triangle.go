package letgo_pascals_triangle

import (
	_ "letgo_repo/system_file/code_enter"
)

/*杨辉三角 | https://leetcode.cn/problems/pascals-triangle*/

/*
给定一个非负整数 numRows，生成「杨辉三角」的前 numRows 行。

在「杨辉三角」中，每个数是它左上方和右上方的数的和。





示例 1:

输入: numRows = 5
输出: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
示例 2:

输入: numRows = 1
输出: [[1]]


提示:

1 <= numRows <= 30

*/

func generate(numRows int) [][]int {

	answer := getInitArr(numRows)

	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			if isBound(i, j) {
				answer[i][j] = 1
				continue
			}

			answer[i][j] = answer[i-1][j-1] + answer[i-1][j] // 排除边界情况下的推导公式
		}
	}

	return answer
}

func isBound(i int, j int) bool {
	if j == 0 {
		return true
	}

	if j == i {
		return true
	}

	return false
}

func getInitArr(numRows int) [][]int {
	answer := make([][]int, numRows)

	for i := 0; i < numRows; i++ {
		answer[i] = make([]int, i+1)
	}

	return answer
}
