package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

//当接收到Control+c,kill -1,kill -2 的时候,都可以执行执行defer函数
// kill -9依然不会正常退出。

func main() {
	fmt.Println("Application is begin.")

	//当程序接受到退出信号的时候,将会执行
	defer fmt.Println("Application is end.")

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("Application is running")
		}
	}()

	// 捕获程序退出信号
	msgChan := make(chan os.Signal, 1)
	signal.Notify(msgChan, os.Kill, os.Interrupt)
	<-msgChan
}
