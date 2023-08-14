package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/mcs-unity/onvif/pkg/wsdiscovery"
)

var FilePath = ""

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	FilePath = wd

	r, err := os.ReadFile(fmt.Sprintf("%s/%s", FilePath, "pkg/wsdiscovery/prob.xml"))
	if err != nil {
		panic(err)
	}

	if _, err := wsdiscovery.SendMulticastProbe(r); err != nil {
		panic(err)
	}

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	fmt.Println("Hello world")
}
