package main

import (
	"fmt"
	"time"
)

func foo() {
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Hello, playground")
	}()

	fmt.Println("Hello foo")
	return
}
