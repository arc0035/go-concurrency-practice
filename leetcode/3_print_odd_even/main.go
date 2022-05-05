package main

import (
	"fmt"
	"sync"
)

var zeroCh = make(chan interface{})
var oddCh = make(chan int)
var evenCh = make(chan int)

func main() {
	fmt.Println("start")
	z := ZeroEvenOdd{
		n: 5,
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		z.Zero()
		wg.Done()
	}()

	go func() {
		z.Odd()
	}()

	go func() {
		z.Even()
	}()

	zeroCh <- nil
	wg.Wait()
}

type ZeroEvenOdd struct {
	n int
}

func (z ZeroEvenOdd) Zero() {
	<-zeroCh
	for i := 1; i <= z.n; i++ {
		fmt.Print("0")
		if (i & 1) == 1 {
			oddCh <- i
		} else {
			evenCh <- i
		}
		<-zeroCh
	}

}

func (z ZeroEvenOdd) Odd() {
	for {
		i := <-oddCh
		fmt.Printf("%d", i)
		zeroCh <- nil
	}

}

func (z ZeroEvenOdd) Even() {
	for {
		i := <-evenCh
		fmt.Printf("%d", i)
		zeroCh <- nil
	}

}
