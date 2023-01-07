package utils

import (
	"fmt"
	"reflect"
)

type ThreadUtilI interface {
	IsAllDone() bool
	IsDone(threadName string) bool
	AddThread(runFunc interface{}) (threadName string)
	CompleteNow(threadName string)
}

var ThreadUtil ThreadUtilI = &ThreadUtilImpl{}

type ThreadUtilImpl struct {
	ThreadMap map[string]bool // true 代表激活状态
}

func (t *ThreadUtilImpl) IsAllDone() bool {
	if len(t.ThreadMap) == 0 {
		return true
	}

	for _, value := range t.ThreadMap {
		if value {
			return false
		}
	}

	return true
}

func (t *ThreadUtilImpl) IsDone(threadName string) bool {
	if len(t.ThreadMap) == 0 {
		return true
	}

	value, ok := t.ThreadMap[threadName]
	if !ok {
		println("无此协程")
		return true
	}

	return !value
}

func (t *ThreadUtilImpl) AddThread(runFunc interface{}) (threadName string) {
	threadName = fmt.Sprintf("thread [%d]", len(t.ThreadMap)+1)
	if t.ThreadMap == nil {
		t.ThreadMap = make(map[string]bool)
	}
	t.ThreadMap[threadName] = true

	go func() {
		reflect.ValueOf(runFunc).Call([]reflect.Value{})

		t.ThreadMap[threadName] = false
	}()

	return threadName
}

func (t *ThreadUtilImpl) CompleteNow(threadName string) {
	if len(t.ThreadMap) == 0 {
		return
	}
	t.ThreadMap[threadName] = false
	fmt.Printf("completed %s\n", threadName)
}
