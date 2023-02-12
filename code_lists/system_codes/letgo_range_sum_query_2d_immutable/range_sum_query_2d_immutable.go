package letgo_range_sum_query_2d_immutable

import (
	_ "letgo_repo/system_file/code_enter"
)

/*二维区域和检索 - 矩阵不可变 | https://leetcode.cn/problems/range-sum-query-2d-immutable*/

type NumMatrix struct {
	Matrix    [][]int
	sumMatrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	var nm NumMatrix
	nm.Matrix = append(nm.Matrix, matrix...)
	nm.sumMatrix = make([][]int, len(nm.Matrix))
	per := 0
	for i, ta := range nm.Matrix {
		for j := range ta {
			if j == 0 {
				continue
			}
			per = ta[j-1]

			nm.Matrix[i][j] = per + nm.Matrix[i][j]
		}
	}
	return nm
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		if col1 > 0 {
			sum += this.Matrix[i][col2] - this.Matrix[i][col1-1]
		} else {
			sum += this.Matrix[i][col2]
		}

	}

	return sum
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
