package server

import (
	"errors"
	"net"
	"net/http"
	"sync"

	"github.com/mcs-unity/onvif/internal/listener"
)

func (s *Server) Stop() error {
	s.l.Lock()
	defer s.l.Unlock()

	port := listener.GetPort(s.con)
	if port == 0 {
		return errors.New("unable to fetch port to close the server")
	}

	return listener.Close(port)
}

func (s *Server) Listen(h http.Handler) error {
	s.l.Lock()
	defer s.l.Unlock()

	if s.tls != nil {
		return http.ServeTLS(s.con, h, s.tls.certPath, s.tls.keyPath)
	}

	return http.Serve(s.con, h)
}

func New(c net.Listener, t tls) IServer {
	return &Server{
		l:   &sync.Mutex{},
		con: c,
		tls: &t,
	}
}

func Cert(certPath, keyPath string) tls {
	return tls{
		certPath: certPath,
		keyPath:  keyPath,
	}
}
