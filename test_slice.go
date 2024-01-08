package main

import "fmt"

func SliceExtend() {
	var slice []int
	fmt.Printf("slice 的长度为[%d],容量为[%d] \n", len(slice), cap(slice))
	s1 := append(slice, 1, 2, 3)
	fmt.Printf("s1 的长度为[%d],容量为[%d] \n", len(s1), cap(s1))
	s2 := append(s1, 4)
	fmt.Printf("s2 的长度为[%d],容量为[%d] \n", len(s2), cap(s2))

	fmt.Println(&s1[0] == &s2[0])

}

func SliceRice(s []int) {
	s = append(s, 0)
	for i := range s {
		s[i]++
	}
}

func SlicePrint() {
	s1 := []int{1, 2}
	s2 := s1
	s2 = append(s2, 3)
	SliceRice(s1)
	SliceRice(s2)
	fmt.Println(s1, s2)
}

func main() {
	SliceExtend()

	SlicePrint()

	testStr := "yes你好"

	fmt.Printf("%s", testStr[3:])

}
