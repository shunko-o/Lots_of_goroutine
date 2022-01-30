package main

import (
	"fmt"
	"sync"
	"time"
)

type CountMap struct {
	value map[string]int
	sync.Mutex
}

func New() *CountMap {
	return &CountMap{
			value: make(map[string]int),
	}
}

func (cm *CountMap) Set(key string) {
	cm.Lock()
	defer cm.Unlock()
	cm.value[key]++
}

func (cm *CountMap) Get(key string) int {
	cm.Lock()
	defer cm.Unlock()
	return cm.value[key]
}

func main() {
	cm := New()

	go func () {
		for i := 0; i < 10; i++ {
			cm.Set("shunko")
		}
	}()

	go func () {
		for i := 0; i < 10; i++ {
			cm.Set("shunko")
		}
	}()

	time.Sleep(time.Second)
	fmt.Println(cm.Get("shunko"))
}
