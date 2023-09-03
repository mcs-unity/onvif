package listener

import (
	"net"
	"sync"
)

type network = string
type port = uint16

const (
	TCP network = "tcp"
)

type hub struct {
	l    sync.Locker
	list map[port]net.Listener
}
