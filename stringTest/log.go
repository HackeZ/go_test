package stringTest

import (
	"log"
)

func LogByStitchString() {
	log.Println("I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I")
	log.Println("I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I")
	log.Println("I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I")
	log.Println("I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I" + "I")
}

func LogByMultiString() {
	log.Println("I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I")
	log.Println("I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I")
	log.Println("I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I")
	log.Println("I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I", "I")
}

func LogByFormatString() {
	str := "I"
	log.Printf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s\n", str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str)
	log.Printf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s\n", str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str)
	log.Printf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s\n", str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str)
	log.Printf("%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s%s\n", str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str, str)
}
