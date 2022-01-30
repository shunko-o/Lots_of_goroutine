package main

import (
	"fmt"
	"time"
)

func goroutine1(ch chan string) {
	for {
		ch <- "Go1"
		time.Sleep(1 * time.Second)
	}
}

func goroutine2(ch chan string) {
	for {
		ch <- "Go2"
		time.Sleep(1 * time.Second)
	}
}

// 永遠に渡し続ける
func main(){
	ch1 := make(chan string)
	ch2 := make(chan string)
	go goroutine1(ch1)
	go goroutine2(ch2)

	for {
		select{
		case s1 := <- ch1:
			fmt.Println(s1)
		case s2 := <- ch2:
			fmt.Println(s2)
		}
	}
}
