package server

import (
	"net"
	"net/http"
	"sync"
	"time"
)

const timeout = 10 * time.Second

type IServer interface {
	Listen(h http.Handler) error
	Terminate() error
}

type TLSConfig struct {
	certPath string
	keyPath  string
}

type Server struct {
	con  net.Listener
	http http.Server
	mu   sync.Locker
	tls  *TLSConfig
}
