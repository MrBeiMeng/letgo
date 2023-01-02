package code_lists

/*最富有客户的资产总量 | https://leetcode.cn/problems/richest-customer-wealth*/

// maximumWealth
//
// @Description: 给你一个 m x n 的整数网格 accounts ，其中 accounts[i][j] 是第 i 位客户在第 j 家银行托管的资产数量。返回最富有客户所拥有的 资产总量 。
// 客户的 资产总量 就是他们在各家银行托管的资产数量之和。最富有客户就是 资产总量 最大的客户。
// @param accounts
// @return int
func maximumWealth(accounts [][]int) int {
	maxV := 0

	for _, propertyArr := range accounts {
		var sum int = SumProperty(propertyArr)
		maxV = Max(sum, maxV)
	}

	return maxV
}

func SumProperty(arr []int) (result int) {

	for _, item := range arr {
		result += item
	}

	return
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
