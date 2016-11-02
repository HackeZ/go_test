package main

import (
	"fmt"
	hessian "github.com/viant/gohessian"
)

type Person struct {
	firstName string
	lastName  string
}

func main() {
	person := Person{"John", "Doe"}
	gh := hessian.NewGoHessian(nil, nil)
	byteArray, _ := gh.ToBytes(person)

	fmt.Println("Person +> ", string(byteArray))
}
