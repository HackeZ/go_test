package main

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"go_test/server-controller_test/svc3/svc"
)

// Program represents a runable program.
type Program struct {
	ctx        context.Context
	exitFunc   context.CancelFunc
	cancelFunc map[string]context.CancelFunc
	wg         WaitGroupWrapper
}

func (p *Program) Init() error {
	// just demo, do something here.
	// like Consul, ZoomKeeper, DB Connect Pool
	return nil
}

func (p *Program) Start() error {
	fmt.Println("本程序将会根据输入,启动或终止服务。")

	reader := bufio.NewReader(os.Stdin)
	go func() {
		for {
			fmt.Println("程序退出命令:exit;服务启动命令:<start||s>-[name];服务停止命令:<cancel||c>-[name]。请注意大小写!")
			input, _, _ := reader.ReadLine()
			command := string(input)

			if command == "exit" {
				goto OutLoop
			} else {
				command, name, err := splitInput(input)
				if err != nil {
					fmt.Println(err)
					continue
				}

				switch command {
				case "start", "s":
					newctx, cancelFunc := context.WithCancel(p.ctx)
					p.cancelFunc[name] = cancelFunc

					p.wg.Wrap(func() {
						Func(newctx, name)
					})
				case "cancal", "c":
					cancelFunc, exist := p.cancelFunc[name]
					if exist {
						cancelFunc()
					}
				}
			}
		}
	OutLoop:
		//由于程序退出被Run的os.Notify阻塞，因此调用以下方法通知退出代码执行。
		svc.Interrupt()
	}()
	return nil
}

func (p *Program) Stop() error {
	p.exitFunc()
	p.wg.Wait()
	fmt.Println("所有服务终止,程序退出!")
	return nil
}

// =====================
// ======= Tools =======
// =====================

func splitInput(input []byte) (command, name string, err error) {
	line := string(input)
	strs := strings.Split(line, "-")
	if strs == nil || len(strs) != 2 {
		err = errors.New("输入不符合规则。")
		return
	}
	command = strs[0]
	name = strs[1]
	return
}

// Func 一个简单的循环方法,模拟被加载、释放的微服务
func Func(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			goto OutLoop
		case <-time.Tick(time.Second * 2):
			fmt.Printf("%s is running.\n", name)
		}
	}
OutLoop:
	fmt.Printf("%s is end.\n", name)
}

// WaitGroupWrapper 封装 WaitGroup 结构
type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(f func()) {
	w.Add(1)
	go func() {
		f()
		w.Done()
	}()
}
