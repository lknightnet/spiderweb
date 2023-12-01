package main

import (
	"fmt"
	"spiderweb/log"
)

func main() {
	l := log.NewLogger("http://localhost:8000/log")

	o, _ := log.NewOther("key", "value")
	msg := log.NewMessage("info", "text", "test service", o)

	err := l.Print(msg)
	if err != nil {
		fmt.Println(err)
	}
}
