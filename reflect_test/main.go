package main

import (
	"fmt"
	"reflect"
)

var ObjectHessian struct {
	Name      string
	Height    int32
	Type      string
	TestSlice []string
	TestMap   map[string]int
}

func init() {
	// init ObjectHessian
	ObjectHessian.Name = "I come from GoHessian"
	ObjectHessian.Height = 718
	ObjectHessian.Type = "3131"

}

func main() {
	objectType := reflect.TypeOf(ObjectHessian)
	fmt.Println("Object info from ObjectHessian => ", objectType)
	for i := 0; i < objectType.NumField(); i++ {
		fmt.Println("type of objectType[", i, "] =>", objectType.Field(i).Type)
	}
	fmt.Println("Len info ObjectHessian =>", objectType.NumField())
	t, exist := objectType.FieldByName("Type")
	if !exist {
		fmt.Println("Type Field not set!")
		return
	}
	fmt.Println("type of Type in ObjectHessian =>", t)

	objectValue := reflect.ValueOf(ObjectHessian)
	fmt.Println("Object info from ObjectHessian => ", objectValue)
	v := objectValue.FieldByName("Type")
	if !v.IsValid() || v.Type().String() != "string" {
		fmt.Println("Type Field not set!")
		return
	}
	fmt.Println("value of Type in ObjectHessian =>", v.Type().String())
	//reflect
}
