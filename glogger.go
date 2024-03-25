package glogger

import "fmt"

type GLogger struct {
	test string
}

func (l *GLogger) print() {
	fmt.Println("hello from GLogger")
}
