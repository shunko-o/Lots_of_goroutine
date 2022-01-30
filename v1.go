package main

import (
	"fmt"
	"sync"
)

func goroutine(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Println(s)
	}
}

func say() {
	fmt.Println("Hello!!!")
}

func main(){
	// WaitGroupを使用しないと、say関数が呼ばれた瞬間にプログラムが終了してしまう。
	var wg sync.WaitGroup
	wg.Add(1)
	go goroutine("call goroutine", &wg)
	say()
	wg.Wait()
}
