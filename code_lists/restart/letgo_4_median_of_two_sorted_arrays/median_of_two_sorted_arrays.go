package letgo_median_of_two_sorted_arrays

import (
	_ "letgo_repo/system_file/code_enter"
	"math"
)

/*寻找两个正序数组的中位数 | https://leetcode.cn/problems/median-of-two-sorted-arrays*/

/*
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

算法的时间复杂度应该为 O(log (m+n)) 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

提示：

nums1.length == m
nums2.length == n
0 <= m <= 1000
0 <= n <= 1000
1 <= m + n <= 2000
-106 <= nums1[i], nums2[i] <= 106

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	// 思路，将左长右短两个数组从划分点分开，保证两边的数组尽量相同，左边可以比右边多一个元素，这时如果满足左边的最大值小于右边的最大值。就完成了任务。否则，如果mid大就排掉mid，mid2大就mid+1

	if len(nums2) > len(nums1) {
		nums1, nums2 = nums2, nums1
	}

	//if len(nums2) == 0 { // 排除有一个数组为空的情况，题中表明不会同时空
	//	return findMedianSortedArray(nums1)
	//}

	//left, right := 0, len(nums1)
	left, right := getMinLeft(nums1, nums2), getMaxRight(nums1, nums2) // 保证mid2划分点不会切出去太多，最多切出去一个

	for left-right <= 1 {

		mid := (right-left)/2 + left
		if right < left {
			mid = right
		}
		mid2 := (len(nums1)+len(nums2)-1)/2 - mid - 1

		//fmt.Printf("%d,%d", mid, mid2)

		var leftHalfMax int = getLeftHalfMax(nums1, mid, nums2, mid2)
		var rightHalfMin int = getRightHalfMin(nums1, mid+1, nums2, mid2+1)

		if leftHalfMax <= rightHalfMin {
			// 完成任务
			if (len(nums1)+len(nums2))%2 == 0 {

				return (float64(leftHalfMax) + float64(rightHalfMin)) / 2
			}

			return float64(leftHalfMax)
		}

		// 如果nums1[mid] > nums2[mid2] right = mid-1 else left = mid+1
		var nums1MidMax int = getLeftMidMax(nums1, mid)
		var nums2Mid2Max int = getLeftMidMax(nums2, mid2)

		if nums1MidMax > nums2Mid2Max {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}

func getLeftMidMax(nums []int, mid int) int {
	answer := math.MinInt

	if indexCorrect(mid, nums) {
		answer = max(answer, nums[mid])
	}

	return answer
}

func getRightHalfMin(nums1 []int, i int, nums2 []int, i2 int) int {
	answer := math.MaxInt

	if indexCorrect(i, nums1) {
		answer = min(nums1[i], answer)
	}

	if indexCorrect(i2, nums2) {
		answer = min(nums2[i2], answer)
	}

	return answer
}

func getLeftHalfMax(nums1 []int, mid int, nums2 []int, mid2 int) int {

	answer := math.MinInt

	if indexCorrect(mid, nums1) {
		answer = max(nums1[mid], answer)
	}

	if indexCorrect(mid2, nums2) {
		answer = max(nums2[mid2], answer)
	}

	return answer
}

func getMaxRight(nums1 []int, nums2 []int) int {
	return (len(nums1) + len(nums2) - 1) / 2
}

func getMinLeft(nums1 []int, nums2 []int) int {
	return (len(nums1) - len(nums2)) / 2
}

func findMedianSortedArray(nums []int) float64 {
	if len(nums)%2 == 0 {
		return (float64(nums[len(nums)/2]) + float64(nums[len(nums)/2-1])) / 2
	}

	return float64(nums[len(nums)/2])
}

func indexCorrect(index int, nums []int) bool {
	return 0 <= index && index < len(nums)
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// binarySearch
//
//	@Description: 二分查找
//	@param nums
//	@param target
//	@return int
func binarySearch(nums []int, target int) int { // 一个简单的二分查找例子
	left, right := 0, len(nums)

	for left <= right {
		mid := (right-left)/2 + left

		if nums[mid] < target {
			left = mid + 1
			continue
		}

		if nums[mid] > target {
			right = mid - 1
			continue
		}

		return mid
	}

	return -1
}
