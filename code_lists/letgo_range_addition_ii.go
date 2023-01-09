package code_lists

/*范围求和 II | https://leetcode.cn/problems/range-addition-ii*/

// maxCount
//
//	@Description:给你一个 m x n 的矩阵 M ，初始化时所有的 0 和一个操作数组 op ，其中 ops[i] = [ai, bi] 意味着当所有的 0 <= x < ai 和 0 <= y < bi 时， M[x][y] 应该加 1。
//
// 在 执行完所有操作后 ，计算并返回 矩阵中最大整数的个数 。
//
//	@param m
//	@param n
//	@param ops
//	@return int
func maxCount(m int, n int, ops [][]int) int {

	// 按x，y来看的话，二维数组是倒过来的，y对应i，x对应的j
	yMin := m
	xMin := n

	for _, yArr := range ops {
		yMin = maxCountMin(yArr[0], yMin)
		xMin = maxCountMin(yArr[1], xMin)
	}

	answer := yMin * xMin
	return answer
}

func maxCountMin(a, b int) int {
	if a <= b {
		return a
	}

	return b
}
