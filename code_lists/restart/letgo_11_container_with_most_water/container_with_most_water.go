package letgo_container_with_most_water

import (
	_ "letgo_repo/system_file/code_enter"
)

/*盛最多水的容器 | https://leetcode.cn/problems/container-with-most-water*/

/*
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器。



示例 1：



输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：

输入：height = [1,1]
输出：1


提示：

n == height.length
2 <= n <= 105
0 <= height[i] <= 104
*/

func maxArea(height []int) int {

	maxCap, left, right := 0, 0, len(height)-1

	for left < right {
		maxCap = max(maxCap, getCap(height[left], height[right], right-left))

		if height[left] < height[right] {
			left++
			continue
		}

		if height[left] > height[right] {
			right--
			continue
		}

		if right-left == 1 { // 两边相等的情况下
			break
		}

		if height[left+1] <= height[right-1] {
			left++
		} else {
			right--
		}

	}

	return maxCap
}

func getCap(height1, height2, length int) int {
	return length * min(height1, height2)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a <= b {
		return b
	}
	return a
}
