package system_codes

/*H 指数 | https://leetcode.cn/problems/h-index*/

func hIndex(citations []int) int {

	//sort.Sort(sort.Reverse(sort.IntSlice(citations)))

	hIndexCountSort(citations)

	h := 0
	for i, citation := range citations {
		tmpH := hIndexMin(i+1, citation) // i+1 表示当前统计的论文数量
		h = hIndexMax(h, tmpH)
	}

	return h
}

func hIndexCountSort(nums []int) {
	if len(nums) == 1 {
		return
	}

	rightBound := -1
	leftBound := len(nums)

	for _, num := range nums {
		rightBound = hIndexMax(rightBound, num)
		leftBound = hIndexMin(leftBound, num)
	}

	sortList := make([]int, rightBound-leftBound+1)

	for _, num := range nums {
		sortList[num-leftBound] += 1
	}

	nums = nums[:0]
	// 正向排序

	// 反向排序
	for i := len(sortList) - 1; i >= 0; i-- {
		times := sortList[i]
		for j := 0; j < times; j++ {
			nums = append(nums, i+leftBound)
		}
	}
}

func hIndexMin(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func hIndexMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
