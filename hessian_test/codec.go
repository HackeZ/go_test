package main

import (
	"fmt"
	"strings"

	gh "github.com/wahahajun/gohessian"
)

const (
	BeforeHessianString  string  = "1plus2=三abc你好吗"
	BeforeHessianInt     int     = 10
	BeforeHessianInt32   int32   = 11
	BeforeHessianInt64   int64   = 12
	BeforeHessianFloat32 float32 = 15.1
	BeforeHessianFloat64 float64 = 15.2
)

func EncodeAndDecodeTest() {
	fmt.Println("================")
	testStringInHessian()
	fmt.Println("================")
	testIntInHessian()
	fmt.Println("================")
	testInt32InHessian()
	fmt.Println("================")
	testInt64InHessian()
	fmt.Println("================")
	testFloat32InHessian()
	fmt.Println("================")
	testFloat64InHessian()
	fmt.Println("================")
}

func testStringInHessian() {
	fmt.Println("Hessian Test For String")
	bytes, err := gh.Encode(BeforeHessianString)
	if err != nil {
		fmt.Println("Hessian Encode Error:", err)
		return
	}
	fmt.Println("Encode => ", bytes)
	r := strings.NewReader(string(bytes))
	decode, err := gh.NewHessian(r).Parse()
	if err != nil {
		fmt.Println("Hessian Decode Error:", err)
		return
	}
	fmt.Println("Decode =>", decode)
	if decode.(string) != BeforeHessianString {
		fmt.Println("Bad Hessian")
		return
	}
	fmt.Println("Success Hessian")
}

func testIntInHessian() {
	fmt.Println("Hessian Test For Int")
	bytes, err := gh.Encode(BeforeHessianInt)
	if err != nil {
		fmt.Println("Hessian Encode Error:", err)
		return
	}
	fmt.Println("Encode => ", bytes)
	r := strings.NewReader(string(bytes))
	decode, err := gh.NewHessian(r).Parse()
	if err != nil {
		fmt.Println("Hessian Decode Error:", err)
		return
	}
	fmt.Println("Decode =>", decode)
	if decode.(int32) != int32(BeforeHessianInt) {
		fmt.Println("Bad Hessian")
		return
	}
	fmt.Println("Success Hessian")
}

func testInt32InHessian() {
	fmt.Println("Hessian Test For Int32")
	bytes, err := gh.Encode(BeforeHessianInt32)
	if err != nil {
		fmt.Println("Hessian Encode Error:", err)
		return
	}
	fmt.Println("Encode => ", bytes)
	r := strings.NewReader(string(bytes))
	decode, err := gh.NewHessian(r).Parse()
	if err != nil {
		fmt.Println("Hessian Decode Error:", err)
		return
	}
	fmt.Println("Decode =>", decode)
	if decode.(int32) != BeforeHessianInt32 {
		fmt.Println("Bad Hessian")
		return
	}
	fmt.Println("Success Hessian")
}

func testInt64InHessian() {
	fmt.Println("Hessian Test For Int64")
	bytes, err := gh.Encode(BeforeHessianInt64)
	if err != nil {
		fmt.Println("Hessian Encode Error:", err)
		return
	}
	fmt.Println("Encode => ", bytes)
	r := strings.NewReader(string(bytes))
	decode, err := gh.NewHessian(r).Parse()
	if err != nil {
		fmt.Println("Hessian Decode Error:", err)
		return
	}
	fmt.Println("Decode =>", decode)
	if decode.(int64) != BeforeHessianInt64 {
		fmt.Println("Bad Hessian")
		return
	}
	fmt.Println("Success Hessian")
}

func testFloat32InHessian() {
	fmt.Println("Hessian Test For Float32")
	bytes, err := gh.Encode(BeforeHessianFloat32)
	if err != nil {
		fmt.Println("Hessian Encode Error:", err)
		return
	}
	fmt.Println("Encode => ", bytes)
	r := strings.NewReader(string(bytes))
	decode, err := gh.NewHessian(r).Parse()
	if err != nil {
		fmt.Println("Hessian Decode Error:", err)
		return
	}
	fmt.Println("Decode =>", decode)
	if decode.(float32) != BeforeHessianFloat32 {
		fmt.Println("Bad Hessian")
		return
	}
	fmt.Println("Success Hessian")
}

func testFloat64InHessian() {
	fmt.Println("Hessian Test For Float64")
	bytes, err := gh.Encode(BeforeHessianFloat64)
	if err != nil {
		fmt.Println("Hessian Encode Error:", err)
		return
	}
	fmt.Println("Encode => ", bytes)
	r := strings.NewReader(string(bytes))
	decode, err := gh.NewHessian(r).Parse()
	if err != nil {
		fmt.Println("Hessian Decode Error:", err)
		return
	}
	fmt.Println("Decode =>", decode)
	if decode.(float64) != BeforeHessianFloat64 {
		fmt.Println("Bad Hessian")
		return
	}
	fmt.Println("Success Hessian")
}
