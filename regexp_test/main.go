package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

var ignoreSlice = flag.String("is", "", "Set Up ignore slice.")

var ignoreList []string

func init() {
	flag.Parse()

	ignoreList = make([]string, 0)

	if len(*ignoreSlice) != 0 {
		for _, v := range strings.Split(*ignoreSlice, ",") {
			ignoreList = append(ignoreList, v)
		}
	}

	fmt.Println("ignore List =>", ignoreList)
}

// RegexpMatch using regexp.MatchString and return result only.
// @param string
// @return bool
func RegexpMatch(FileName string) (match bool) {
	for _, v := range ignoreList {
		match, _ := regexp.MatchString(v, FileName)
		fmt.Println("match =>", match)
		if match == true {
			return match
		}
	}
	return
}

func main() {
	match := RegexpMatch("test.js")
	fmt.Println("test.js => ", match)
}
