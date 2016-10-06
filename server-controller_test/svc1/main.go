package main

import (
	"fmt"
	"time"
)

//当接收到Control+c,kill -1,kill -2,kill -9 均无法正常执行defer函数
func main() {
	fmt.Println("Application is begin.")

	// 以下代码不会执行
	defer fmt.Println("Application is end.")

	for {
		time.Sleep(time.Second)
		fmt.Println("Application is running.")
	}
}
