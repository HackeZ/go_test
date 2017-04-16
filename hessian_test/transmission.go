package main

import (
	"fmt"

	//"bytes"

	gh "github.com/hackez/gohessian"
)

const (
	HessianCPPServiceHost = "http://hessian.caucho.com"
	HessianCPPServiceURL  = "/test/test2"
)

func TransmissionTest() {
	//fmt.Println("================")
	// testToConnOfficalService("argTypedMap_1")
	//testToConnOfficalService("replyObject_2")
	fmt.Println("================")
	//testToConnUserSpaceService("queryMemberList")
	testToConnUserSpaceService()
	//testToConnUserSpaceService("veryfyPwdQuestion")

}

func testToConnOfficalService(method string) {

}

func testToConnUserSpaceService() {
}
