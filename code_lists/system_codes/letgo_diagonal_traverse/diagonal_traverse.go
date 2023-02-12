package letgo_diagonal_traverse

import (
	_ "letgo_repo/system_file/code_enter"
)

/*对角线遍历 | https://leetcode.cn/problems/diagonal-traverse*/

func findDiagonalOrder(mat [][]int) []int {
	ansArr := make([]int, 0)
	deep := 1
	row, col := 0, 0
	for true {
		if row > len(mat)-1 || col > len(mat[0])-1 {
			break
		}

		ansArr = append(ansArr, mat[row][col])
		if deep%2 != 0 {
			if row == 0 || col == len(mat[0])-1 {
				if col == len(mat[0])-1 {
					row += 1
					col -= 1
				}

				col += 1
				deep++
				continue
			}

			row -= 1
			col += 1
		} else {
			if col == 0 || row == len(mat)-1 {
				if row == len(mat)-1 {
					col += 1
					row -= 1
				}
				row += 1
				deep++
				continue
			}

			row += 1
			col -= 1
		}

	}

	return ansArr
}
