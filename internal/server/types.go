package server

import (
	"net"
	"net/http"
	"sync"
)

type IServer interface {
	Listen(h http.Handler) error
	Stop() error
}

type tls struct {
	certPath string
	keyPath  string
}

type Server struct {
	con net.Listener
	mu  sync.Locker
	tls *tls
}
