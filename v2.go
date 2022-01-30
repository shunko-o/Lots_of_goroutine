package main

import (
	"fmt"
)

func goroutine1(num []int, c chan int) {
	count := 0
	for _, v := range num {
		count += v
	}
	c <- count
}

func goroutine2(num []int, c chan int) {
	count := 0
	for _, v := range num {
		count += v
	}
	c <- count
}

func main() {
	i5 := []int{1, 2, 3, 4, 5}
	i10 := []int{1,2,3,4,5,6,7,8,9,10}
	c := make(chan int)

	go goroutine1(i5, c)
	fmt.Println(<-c)
	go goroutine2(i10, c)
	fmt.Println(<-c)
}
