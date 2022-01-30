package main

import "fmt"


func goroutine1(ch1 chan int){
	defer close(ch1)
	for i:=0; i<10; i++ {
		ch1 <- i
	}
}

// firstは受信のチャネル、secondは送信のチャネル。以下のように明示的に定義することも出来る
// first <-chan int, second chan<- int
func goroutine2(ch1 chan int, ch2 chan int){
	defer close(ch2)
	for i := range ch1 {
		ch2 <- i * 10
	}
}

func goroutine3(ch2 chan int, ch3 chan int){
	defer close(ch3)
	for i := range ch2 {
		ch3 <- i * 100
	}
}

func main(){
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go goroutine1(ch1)
	go goroutine2(ch1, ch2)
	go goroutine3(ch2, ch3)

	for result := range ch3 {
		fmt.Println(result)
	}
}
