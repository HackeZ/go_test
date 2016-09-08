package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"log"
// 	"os"
// )

import (
	"github.com/astaxie/beego"
)

// func main() {
// 	file, err := os.OpenFile("./file.txt", os.O_RDONLY, os.ModePerm)
// 	defer file.Close()

// 	if err != nil {
// 		log.Println("file.txt open false!")
// 	}

// 	fileReader := bufio.NewReader(file)

// 	for {
// 		line, err := fileReader.ReadString('\n')
// 		if err != nil {
// 			if err == io.EOF {
// 				break
// 			}
// 			log.Println("Read File Content Failed!", err.Error())
// 			return
// 		}
// 		fmt.Println(line)
// 	}

// 	fmt.Println("File Read Done!")
// }

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}
