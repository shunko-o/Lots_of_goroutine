package main

import (
	"fmt"
	"sync"
)

func goroutine1(ch chan int, i int) {
	ch <- i * 2
}

func goroutine2(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Println(i)
		wg.Done()
	}
}

func main(){
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go goroutine1(ch, i)
	}
	defer close(ch)

	go goroutine2(ch, &wg)
	wg.Wait()
}
