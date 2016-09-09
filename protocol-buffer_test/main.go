package main

import (
	"fmt"
	"go_test/protocol-buffer_test/human"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	man := &human.Person{
		Id:    1234,
		Name:  "HackerZ",
		Email: "hackerzgz@gmail.com",
		Phones: []*human.Person_PhoneNumber{
			{Number: "555-4321", Type: human.Person_HOME},
			{Number: "1803548271", Type: human.Person_MOBILE},
		},
	}

	// Write the Human back to disk.
	out, err := proto.Marshal(man)
	if err != nil {
		log.Fatalln("Failed to encode Human:", err)
	}
	if err := ioutil.WriteFile("man.txt", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	// Read the existing address book.
	in, err := ioutil.ReadFile("man.txt")
	if err != nil {
		log.Fatalln("Error reading file:", err)
	}
	man2 := &human.Person{}
	if err := proto.Unmarshal(in, man2); err != nil {
		log.Fatalln("Failed to parse man2:", err)
	}
	fmt.Println("man2 -->", man2)
}
