package listener

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync"
)

var list hub

func Get(p port) net.Listener {
	list.l.Lock()
	defer list.l.Unlock()
	if v, ok := list.list[p]; ok {
		return v
	}

	return nil
}

func GetList(p ...port) []net.Listener {
	cons := []net.Listener{}
	for _, port := range p {
		if l := Get(port); l != nil {
			cons = append(cons, l)
		}
	}
	return cons
}

func GetPort(l net.Listener) uint16 {
	sp := strings.Split(l.Addr().String(), ":")
	if len(sp) < 2 {
		return 0
	}
	port, err := strconv.ParseUint(sp[1], 0, 16)
	if err != nil {
		return 0
	}
	return uint16(port)

}

func Listener(ip string, p port, n network) (net.Listener, error) {
	list.l.Lock()
	defer list.l.Unlock()

	addr := net.ParseIP(ip)
	if addr == nil {
		return nil, errors.New("unable to parse ip address")
	}

	l, err := net.ListenTCP(n, &net.TCPAddr{
		IP:   addr,
		Port: int(p),
	})

	if err != nil {
		return nil, err
	}

	if _, ok := list.list[p]; ok {
		return nil, fmt.Errorf("there is already a connection listening on port: %d", p)
	}

	list.list[p] = l
	return l, nil
}

func Close(p port) error {
	if l := Get(p); l != nil {
		if err := l.Close(); err == nil {
			delete(list.list, p)
			return nil
		} else {
			return err
		}
	}

	return fmt.Errorf("no listener listening on port: %d", p)
}

func init() {
	list = hub{
		l:    &sync.Mutex{},
		list: make(map[uint16]net.Listener),
	}
}
