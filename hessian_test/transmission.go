package main

import (
	"fmt"

	gh "github.com/wahahajun/gohessian"
)

const (
	HessianCPPServiceHost = "http://hessian.caucho.com"
	HessianCPPServiceURL  = "/test/test2"
)

func TransmissionTest() {
	fmt.Println("================")
	testToConnOfficalService()
}

func testToConnOfficalService() {
	testAddr := HessianCPPServiceHost + HessianCPPServiceURL
	res, err := gh.Request(testAddr, "replyString_1024")
	if err != nil {
		fmt.Println("Connect To Hessian Offical Service Failed:", err)
		return
	}
	fmt.Println("Result of ", testAddr, " was => ", res)
}
