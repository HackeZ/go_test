package svc

import (
	"os"
	"os/signal"
)

// Service 标准程序执行和退出的执行接口,运行程序要实现接口定义的方法
type Service interface {
	// Init
	Init() error
	// Start 程序启动运行的时候,需要执行的代码。不得阻塞。
	Start() error
	// Stop 程序退出的时候,需要执行的代码。不得阻塞。
	Stop() error
}

var msgMes = make(chan os.Signal, 1)

// Run 程序运行、退出的包装容器,主程序直接调用。
func Run(service Service) error {
	if err := service.Init(); err != nil {
		return err
	}
	if err := service.Start(); err != nil {
		return err
	}

	signal.Notify(msgMes, os.Interrupt, os.Kill)
	<-msgMes
	return service.Stop()
}

// 通常不需要调用,特殊情况下,在程序内其他模块中，需要通知程序退出才会使用。
func Interrupt() {
	msgMes <- os.Interrupt
}
