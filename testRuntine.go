package main

import (
	"errors"
	"sync"
	"time"
)

var wg sync.WaitGroup

func wait3s() {
	time.Sleep(3 * time.Second)
	println("三秒已过")
	wg.Done()
}

func wait5s() {
	time.Sleep(5 * time.Second)
	println("五秒已过")
	wg.Done()
}

func testFunc() error {
	b := false
	if !b {
		return errors.New("err")
	}

	return nil

}

func main() {
	println("开始等待1s")
	time.Sleep(time.Second)
	println("同时开始执行等待五秒、等待三秒函数")
	wg.Add(2)
	go wait5s()
	go wait3s()

	go func() {
		err := testFunc()
		if err != nil {

		}
	}()

	wg.Wait()
	println("协程都已运行完毕")
}
