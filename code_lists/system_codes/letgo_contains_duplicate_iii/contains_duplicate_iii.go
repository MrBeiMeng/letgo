package letgo_contains_duplicate_iii

import (
	_ "letgo_repo/system_file/code_enter"
)

/*存在重复元素 III | https://leetcode.cn/problems/contains-duplicate-iii*/

/*
给你一个整数数组 nums 和两个整数 k 和 t 。请你判断是否存在 两个不同下标 i 和 j，使得 abs(nums[i] - nums[j]) <= t ，同时又满足 abs(i - j) <= k 。

如果存在则返回 true，不存在返回 false。
*/
func containsNearbyAlmostDuplicate(nums []int, indexDiff int, valueDiff int) bool {
	leftIndex := 0
	var maxHeap MaxHeap

	for rightIndex, num := range nums {
		if rightIndex-leftIndex > indexDiff {
			maxHeap.Remove(nums[leftIndex])
			leftIndex++
		}

		if maxHeap.ExistBetween(num-valueDiff, num+valueDiff) {
			return true
		}

		maxHeap.Push(num)
	}

	return false

}

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}

	return num
}

type MaxHeap []int

func (m *MaxHeap) ExistBetween(x, y int) bool {

	// 查找在这个范围内是否有值，前提 x <= y
	if len(*m) == 0 {
		return false
	}

	if (*m)[0] > y || (*m)[len(*m)-1] < x {
		return false
	}

	leftIndex, rightIndex := 0, len(*m)-1
	for leftIndex <= rightIndex && rightIndex >= 0 && leftIndex < len(*m) {
		midIndex := leftIndex + (rightIndex-leftIndex)/2

		if (*m)[midIndex] >= x && (*m)[midIndex] <= y {
			return true
		}

		if (*m)[midIndex] < x {
			leftIndex = midIndex + 1
		}

		if (*m)[midIndex] > y {
			rightIndex = midIndex - 1
		}

	}

	return false
}

func (m *MaxHeap) Len() int {
	return len(*m)
}

func (m *MaxHeap) Remove(x int) {
	for index, num := range *m {
		if num == x {
			*m = append((*m)[:index], (*m)[index+1:]...) // 删除中间1个元素
		}
	}
}

func (m *MaxHeap) Push(x int) {
	// m 要保持有序的状态
	// 二分插入排序
	// 查找在这个范围内是否有值，前提 x <= y

	if len(*m) == 0 {
		*m = append(*m, x)
		return
	}

	if (*m)[0] > x {
		*m = append([]int{x}, (*m)[0:]...)
		return
	}

	if (*m)[len(*m)-1] < x {
		*m = append(*m, x)
		return
	}

	leftIndex, rightIndex := 0, len(*m)-1
	for leftIndex <= rightIndex && rightIndex > 0 && leftIndex < len(*m) {
		midIndex := leftIndex + (rightIndex-leftIndex)/2

		if rightIndex-leftIndex <= 30 {
			// 开始查找元素
			for i := 0 + leftIndex; i <= rightIndex; i++ {
				if (*m)[i] >= x {
					// 在这里添加
					*m = append((*m)[:i], append([]int{x}, (*m)[i:]...)...) // 在i处添加一个元素
					return
				}
			}
		}

		if (*m)[midIndex] < x {
			leftIndex = midIndex + 1
		}

		if (*m)[midIndex] > x {
			rightIndex = midIndex - 1
		}

	}

}

func (m *MaxHeap) Pop() int {
	answer := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]

	return answer
}
