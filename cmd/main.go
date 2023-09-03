package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/mcs-unity/onvif/internal/env"
	"github.com/mcs-unity/onvif/internal/listener"
	"github.com/mcs-unity/onvif/internal/server"
)

func main() {
	c := make(chan os.Signal, 1)
	if err := env.LoadVariables(".env"); err != nil {
		panic(err)
	}
	fmt.Println("env is set to:", os.Getenv("ENV"))
	p, err := strconv.ParseUint(os.Getenv("PORT"), 0, 16)
	if err != nil {
		panic(err)
	}
	port := uint16(p)
	l, err := listener.Listener(os.Getenv("IP"), port, listener.TCP)
	if err != nil {
		panic(err)
	}
	defer listener.Close(port)

	cert := server.Cert("cert/server.crt", "cert/server.key")
	s := server.New(l, &cert)
	go func() {
		err := s.Listen(nil)
		fmt.Println(err)
	}()

	signal.Notify(c, syscall.SIGABRT, syscall.SIGINT, syscall.SIGTERM)

	if v, ok := <-c; ok {
		fmt.Println("exit reason: ", v)
	}
}
