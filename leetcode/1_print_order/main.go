package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	nums := []int{3, 2, 1}

	aToB := make(chan struct{})
	bToC := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(3)
	//启动线程A
	go func() {
		WaitingToStart(0, nums)

		fmt.Println("First")
		aToB <- struct{}{}

		wg.Done()
	}()

	//启动线程B
	go func() {
		WaitingToStart(1, nums)

		//Do stuff here
		<-aToB
		fmt.Println("Second")
		bToC <- struct{}{}
		wg.Done()
	}()

	//启动线程C
	go func() {
		WaitingToStart(2, nums)

		//Do stuff here
		<-bToC
		fmt.Println("Third")
		wg.Done()
	}()
	wg.Wait()
}

// WaitingToStart 排在第n位的线程，等待第n-1位的线程执行完毕。
func WaitingToStart(index int, nums []int) {
	d := time.Duration(nums[index]*2) * time.Second
	time.Sleep(d)
}
