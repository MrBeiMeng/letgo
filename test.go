package main

import (
	"reflect"
	"strings"
)

type Person struct {
}

func bala(nums []int, target int) {
	println("yes")
}

func (p Person) GetFunc() interface{} {
	return bala
}

func main() {
	var person Person
	t := reflect.TypeOf(person.GetFunc())

	argsStr := "[1,2],3"
	argsSlice := strings.Split(argsStr, ",")

	if len(argsSlice) != t.NumIn() {
		println("参数数量错误")
		return
	}

	//var args []reflect.Value
	for i := 0; i < t.NumIn(); i++ {
		switch t.In(i).Kind() {
		case reflect.Slice:
			println("slice")
		case reflect.Int:
			println("int")
		}
	}

}
