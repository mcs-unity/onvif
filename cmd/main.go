package main

import (
	"fmt"
	"os"
	"os/signal"
)

var FilePath = ""

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	FilePath = wd

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("Hello world")
}
