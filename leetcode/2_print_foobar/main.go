package main

import (
	"fmt"
	"sync"
)

var ch1 = make(chan interface{}, 1)
var ch2 = make(chan interface{}, 1)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	n := 5
	ch1 <- nil
	fmt.Println("通知线程1启动!")

	go func() {
		PrintFoo(n)
		wg.Done()
	}()

	go func() {
		PrintBar(n)
		wg.Done()
	}()

	wg.Wait()
}

func PrintFoo(n int) {
	for i := 0; i < n; i++ {
		<-ch1
		fmt.Print("foo")
		ch2 <- nil
	}
}

func PrintBar(n int) {
	for i := 0; i < n; i++ {
		<-ch2
		fmt.Print("bar")
		ch1 <- nil
	}
}
