package server

import (
	"errors"
	"fmt"
	"net"
	"sync"
)

var list hub

func listener(ip string, p port) (net.Listener, error) {
	addr := net.ParseIP(ip)
	if addr == nil {
		return nil, errors.New("unable to parse ip address")
	}
	l, err := net.ListenTCP(TCP, &net.TCPAddr{
		IP:   addr,
		Port: int(p),
	})

	if err != nil {
		panic(err)
	}

	return l, nil
}

func Close(p port) error {
	list.l.Lock()
	defer list.l.Unlock()
	if v, ok := list.list[p]; ok {
		return v.Close()
	}
	return fmt.Errorf("no listener listening on port: %d", p)
}

func Listen(ip string, p port) error {
	list.l.Lock()
	defer list.l.Unlock()
	if _, ok := list.list[p]; ok {
		return fmt.Errorf("there is already a connection listening on port: %d", p)
	}

	l, err := listener("", 8080)
	if err != nil {
		return err
	}

	list.list[p] = l
	return nil
}

func init() {
	list = hub{
		l:    &sync.Mutex{},
		list: make(map[uint16]net.Listener),
	}
}
