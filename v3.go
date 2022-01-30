package main

import (
	"fmt"
)

func goroutine(nums []int, ch chan int){
	count := 0
	for _, v := range nums{
		count += v
		ch <- count
	}
	close(ch)
}

func main(){
	nums := []int{1,2,3,4,5}
	ch := make(chan int, len(nums))
	go goroutine(nums, ch)

	for i := range ch {
		fmt.Println(i)
	}
}
