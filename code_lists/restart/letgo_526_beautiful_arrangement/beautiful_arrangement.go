package letgo_beautiful_arrangement

import (
	_ "letgo_repo/system_file/code_enter"
)

/*优美的排列 | https://leetcode.cn/problems/beautiful-arrangement*/

/*
假设有从 1 到 n 的 n 个整数。用这些整数构造一个数组 perm（下标从 1 开始），只要满足下述条件 之一 ，该数组就是一个 优美的排列 ：

perm[i] 能够被 i 整除
i 能够被 perm[i] 整除
给你一个整数 n ，返回可以构造的 优美排列 的 数量 。



示例 1：

输入：n = 2
输出：2
解释：
第 1 个优美的排列是 [1,2]：
    - perm[1] = 1 能被 i = 1 整除
    - perm[2] = 2 能被 i = 2 整除
第 2 个优美的排列是 [2,1]:
    - perm[1] = 2 能被 i = 1 整除
    - i = 2 能被 perm[2] = 1 整除
示例 2：

输入：n = 1
输出：1


提示：

1 <= n <= 15

*/

func countArrangement(n int) int {
	// 回溯算法
	// 每次选一个数字，然后往下递归
	chosen := make([]bool, n)
	count := 0

	backtrack(chosen, 0, &count)

	return count
}

// backtrack
//
//	@Description: 回溯算法寻找完美排列个数
//	@param chosen
//	@param i 表示当前层级，应满足 当前元素 能被 i 整除 或者 i 能被 当前元素 整除, !! 从0开始
//	@param count
func backtrack(chosen []bool, i int, count *int) {
	// 每次选一个合适的数字，然后往下递归

	// 如果所有的数字都被选择过了，或者没有数字可以再选择了，count+1 表示成功一次

	for index, yes := range chosen {
		if yes {
			continue
		}

		if !IsPrefect(index, i) {
			continue
		}

		chosen[index] = true

		if !allChosen(chosen) { // 所有的数字都被用完了
			backtrack(chosen, i+1, count)
		} else {
			*count++ // 记录当前
		}

		chosen[index] = false
	}
}

func allChosen(chosen []bool) bool {

	for _, item := range chosen {
		if !item {
			return false
		}
	}

	return true
}

func IsPrefect(index int, i int) bool {
	return (index+1)%(i+1) == 0 || (i+1)%(index+1) == 0
}
