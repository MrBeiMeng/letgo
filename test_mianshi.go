package main

import (
	"fmt"
)

type A struct {
	Name string
}

type B struct {
	Name string
}

func main() {
	a := A{
		Name: "yes",
	}

	b := A{
		Name: "yes",
	}

	var ia, ib interface{}

	ia = a
	ib = b

	boolEqual := ia == ib

	fmt.Println(boolEqual)

}
