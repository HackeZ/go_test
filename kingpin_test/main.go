package main

import kingpin "gopkg.in/alecthomas/kingpin.v2"

var (
	igFile          = kingpin.Flag("ignoreFile", "set up ignore file list. split by ','").Short('f').Default("").String()
	igDir           = kingpin.Flag("ignoreDir", "set up ignore directory list. split by ','").Short('d').Default("").String()
	dirRoot         = kingpin.Flag("root", "set up directory root.").Short('r').Required().String()
	maxGoroutineNum = kingpin.Flag("maxG", "set up max of running goroutine number.").Short('g').Default("2048").Int64()
)

func init() {
	kingpin.Version("0.2.4")
	kingpin.Parse()
}

func main() {
	// fmt.Println("It is Verbose =>", *verbose)
	// fmt.Printf("Hello %s!\n", *name)
}
