package main

import (
	"bytes"
	"fmt"
	// "reflect"

	gh "github.com/viant/gohessian"
	gh
)

type QuestionSpecDto struct {
	QuestionID   int64
	QuestionName string
}

type Person struct {
	firstName string
	lastName  string
}

func main() {
	var qs QuestionSpecDto
	qs.QuestionID = int64(10010)
	qs.QuestionName = "你妈妈的名字是"

	br := bytes.NewBuffer(nil)
	encoder := gh.NewEncoder(br, nil)
	encoder.WriteObject(qs)
	encoder.RegisterNameType("__type__", "com.pay.pe.dto.QuestionSpecDto")

	fmt.Println("encoder:", br)

	// p := Person{"John", "Doe"}
	gh := gh.NewGoHessian(nil, nil)
	bt, _ := gh.ToBytes(qs)
	// p_new, _ := gh.ToObject(bt)
	fmt.Println("bt", string(bt))

	fmt.Println(string(bt) == br.String())
	// reflect.DeepEqual(p, p_new)
}
