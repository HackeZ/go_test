package main

import (
	"fmt"
	"time"

	"go_test/server-controller_test/svc3/svc"
)

// Program represents a runable program.
type Program struct{}

// Start When Program is Start running.
func (p *Program) Start() error {
	fmt.Println("Application is begin.")
	//必须非阻塞，因此通过协程封装。
	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("Application is running.")
		}
	}()
	return nil
}

// Init When Program is Init
func (p *Program) Init() error {
	//just demon,do nothing
	// like Consul, ZoomKeeper, DB Connect Pool
	return nil
}

// Stop When Program is Stop
func (p *Program) Stop() error {
	fmt.Println("Application is stop.")
	return nil
}

func main() {
	p := &Program{}
	svc.Run()
}
