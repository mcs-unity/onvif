package server

import (
	"fmt"
	"sync"
	"testing"

	"github.com/mcs-unity/onvif/internal/listener"
)

func TestListen(t *testing.T) {
	l, err := listener.Listen("127.0.0.1", 8080, listener.TCP)
	if err != nil {
		t.Error(err)
	}

	s := New(l, nil)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		if err := s.Listen(nil); err != nil {
			fmt.Println(err)
		}
		wg.Done()
	}()

	if err := s.Terminate(); err != nil {
		t.Error(err)
	}

	wg.Wait()
}
