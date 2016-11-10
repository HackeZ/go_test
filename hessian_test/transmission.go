package main

import (
	"fmt"

	//"bytes"

	gh "github.com/hackez/gohessian"
)

const (
	HessianCPPServiceHost = "http://hessian.caucho.com"
	HessianCPPServiceURL  = "/test/test2"

	// Froad UserSpaceService
	UserSpaceServiceHost = "http://10.43.1.127:9200"
	UserSpaceServiceURL  = "/pe_engine/service/userSpecService"
	// UserSpaceServiceURL = "/pe_engine/service/acctSpecService"
)

func TransmissionTest() {
	//fmt.Println("================")
	// testToConnOfficalService("argTypedMap_1")
	fmt.Println("================")
	//testToConnUserSpaceService("queryMemberList")
	testToConnUserSpaceService("verifyToken")
	//testToConnUserSpaceService("veryfyPwdQuestion")

}

func testToConnOfficalService(method string) {
	userSpaceClient := gh.NewClient(HessianCPPServiceHost, HessianCPPServiceURL)
	argMap := make(map[gh.Any]gh.Any, 0)
	argMap["test"] = "hello"
	res, err := userSpaceClient.Invoke(method, argMap)
	if err != nil {
		fmt.Println("Connect To Hessian Offical Service Failed:", err)
		return
	}
	fmt.Printf("result of %s of method %s was => \n%v\n", HessianCPPServiceHost+HessianCPPServiceURL, method, res)
}

func testToConnUserSpaceService(method string) {
	//memberCodes := []gh.Any{}
	//memberCodes = append(memberCodes, 52002018569, 52002017430, 52001970757)
	memberCode := int64(52002018569)
	token := "2dsfbiuh123n1bu123i"

	// MemQS := make(map[gh.Any]gh.Any)
	// MemQS["questionID"] = 33352
	// MemQS["answer"] = "Hessian接口测试"
	// MemQS["__type__"] = "com.pay.pe.dto.QuestionSpecDto"

	//MemQS := []gh.Any{}
	//MemQS = append(MemQS, 33352, "Hessian接口测试", "com.pay.pe.dto.QuestionSpecDto")
	//ReqList := make([]gh.Any, 0)
	// MemQSList := make([]gh.Any, 0)
	// MemQSList = append(MemQSList, MemQS)
	//ReqList = append(ReqList, MemQSList)
	// log.Println("MemQSList:", MemQSList)

	userSpaceClient := gh.NewClient(UserSpaceServiceHost, UserSpaceServiceURL)
	testAddr := UserSpaceServiceHost + UserSpaceServiceURL
	fmt.Println("testAddress:", testAddr, " memberCode:", memberCode)
	res, err := userSpaceClient.Invoke(method, token, memberCode)

	if err != nil {
		fmt.Println("Connect To Hessian Offical Service Failed:", err)
		return
	}
	resMap := res.(map[gh.Any]gh.Any)
	fmt.Println("reult of ", testAddr, " of method ", method, "was => \n", resMap)
}
