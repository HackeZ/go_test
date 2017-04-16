package main

import (
	"bytes"
	"fmt"

	"strings"

	"os"
	"strconv"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	buffer = kingpin.Flag("buffer", "typing buffer in here, example: \"97, 50\"").Short('b').Required().String()

	version = "0.1.0"
)

func init() {
	kingpin.Version(version)
	kingpin.Parse()
}

func main() {
	var buf bytes.Buffer
	var err error
	var i int
	afterTrim := strings.Replace(*buffer, " ", "", -1)
	strSlice := strings.Split(afterTrim, ",")

	for _, s := range strSlice {
		i, err = strconv.Atoi(s)
		if err != nil {
			fmt.Println("Parse String to Int Failed, error:", err)
			os.Exit(-1)
		}
		buf.WriteByte(byte(i))
	}

	fmt.Println("string =>", buf.String())
}
