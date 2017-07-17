package log

import "log"

func Error(v ...interface{}) {
	log.Println("[Error] ", v)
}

func Info(v ...interface{}) {
	log.Println("[Info] ", v)
}

func Debug(v ...interface{}) {
	log.Println("[Debug] ", v)
}
