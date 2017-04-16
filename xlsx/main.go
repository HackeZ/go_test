package main

import (
	"log"

	"github.com/tealeg/xlsx"
)

func main() {
	file, err := xlsx.OpenFile("./test.xlsx")
	if err != nil {
		log.Fatalln("can not open xlsx:", err)
		return
	}
	for sindex, sheet := range file.Sheets {
		if sindex > 0 {
			break
		}
		log.Println(sindex, ":", sheet.Name)
		for rindex, row := range sheet.Rows {
			if 0 == rindex {
				continue
			}
			log.Println("row:", rindex, ":", *row)
			for cindex, cell := range row.Cells {
				log.Print("cell:", cindex, ":", cell.Value)
			}
		}
	}
}
