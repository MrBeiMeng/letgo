package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	randNums := getRandSlice(100)
	fmt.Println(randNums)

	// 插入排序
	insertSort(randNums)

	// 选择排序
	selectSort(randNums)

	// 冒泡排序
	bubbleSort(randNums)
}

func bubbleSort(randNums []int) {
	randNums = getCopied(randNums)
	for i := 0; i < len(randNums)-1; i++ {
		for j := 0; j < len(randNums)-1; j++ {
			if randNums[j] > randNums[j+1] {
				randNums[j], randNums[j+1] = randNums[j+1], randNums[j]
			}
		}
	}
	fmt.Printf("%v\n", randNums)
}

func selectSort(randNums []int) {
	randNums = getCopied(randNums)
	for i := 0; i < len(randNums)-1; i++ {
		smallestIndex := i
		for j := i + 1; j < len(randNums); j++ {
			if randNums[smallestIndex] > randNums[j] {
				smallestIndex = j
			}
		}

		randNums[i], randNums[smallestIndex] = randNums[smallestIndex], randNums[i]
	}
	fmt.Printf("%v\n", randNums)
}

func getRandSlice(length int) []int {
	var randNums []int = make([]int, 0, 100)
	rand.Seed(time.Now().Unix())
	for i := 0; i < length; i++ {
		randNums = append(randNums, rand.Intn(100))
	}
	return randNums
}

func insertSort(randNums []int) {

	randNums = getCopied(randNums)

	for i := 0; i < len(randNums); i++ {
		for j := 0; j < i; j++ {
			if randNums[j] > randNums[i] {
				randNums[j], randNums[i] = randNums[i], randNums[j]
			}
		}
	}
	fmt.Printf("%v\n", randNums)
}

func getCopied(randNums []int) []int {
	tmp := make([]int, len(randNums))
	copy(tmp, randNums)

	randNums = tmp
	return randNums
}
