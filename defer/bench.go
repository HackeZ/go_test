package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//wg()

	Channel()
}

func Channel() string {
	c := make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		printNum(1)
		c <- 1
	}()

	go func() {
		time.Sleep(5 * time.Second)
		printNum(2)
		c <- 2
	}()

	for {
		msg := <-c
		if 2 == msg {
			break
		}
	}

	fmt.Println("channel all finish")
	return ""
}

func printNum(num int) {
	fmt.Println(num)
}

func WG() {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		defer show(1, &wg)
		time.Sleep(5 * time.Second)
	}()

	go func() {
		defer show(2, &wg)
		time.Sleep(1 * time.Second)
	}()
	wg.Wait()

	fmt.Println("wait group all finish")
}

func show(num int, wg *sync.WaitGroup) {
	fmt.Println(num)
	wg.Done()
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
