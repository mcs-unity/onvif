package server

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync"

	"github.com/mcs-unity/onvif/internal/listener"
)

func (s *Server) Terminate() error {
	port := listener.GetPort(s.con)
	if port == 0 {
		return errors.New("unable to fetch port to close the server")
	}
	fmt.Printf("terminating http server running on port: %d", port)
	return listener.Close(port)
}

func (s *Server) Listen(h http.Handler) error {
	fmt.Println("http server is now linking with net listener")
	s.http = http.Server{
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
		Handler:      h,
	}
	if s.tls != nil {
		return s.http.ServeTLS(s.con, s.tls.certPath, s.tls.keyPath)
	}

	return s.http.Serve(s.con)
}

func New(c net.Listener, t *TLSConfig) IServer {
	return &Server{
		mu:  &sync.Mutex{},
		con: c,
		tls: t,
	}
}

func Cert(certPath, keyPath string) TLSConfig {
	return TLSConfig{
		certPath: certPath,
		keyPath:  keyPath,
	}
}
