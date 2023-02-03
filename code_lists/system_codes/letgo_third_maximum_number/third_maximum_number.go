package letgo_third_maximum_number

import (
	"container/heap"
	_ "letgo_repo/system_file/code_enter"
)

/*第三大的数 | https://leetcode.cn/problems/third-maximum-number*/

/*给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。*/

func thirdMax(nums []int) int {
	var bigHeap Heap
	heap.Init(&bigHeap)

	for _, num := range nums {
		heap.Push(&bigHeap, num)

		if bigHeap.Len() > 3 {
			heap.Pop(&bigHeap)
		}
	}

	if bigHeap.Len() == 3 {
		return heap.Pop(&bigHeap).(int)
	}

	for i := 0; i < bigHeap.Len()-1; i++ {
		heap.Pop(&bigHeap)
	}

	return heap.Pop(&bigHeap).(int)
}

type Heap []int

func (s *Heap) Len() int {
	return len(*s)
}

func (s *Heap) Less(i, j int) bool {
	return (*s)[i] < (*s)[j]
}

func (s *Heap) Swap(i, j int) {
	(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
}

func (s *Heap) Push(x interface{}) {

	for _, num := range *s {
		if num == x.(int) {
			return
		}
	}

	*s = append(*s, x.(int))
}

func (s *Heap) Pop() interface{} {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}
