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
	mu   sync.Locker
	list map[port]net.Listener
}
