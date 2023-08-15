package wsdiscovery

import (
	"net"
	"testing"
)

func TestSendMulticastProbe(t *testing.T) {
	b := []byte("hello")
	go func() {
		_, err := SendMulticastProbe(b)
		if err != nil {
			t.Error(err)
		}
	}()
	con, err := net.ListenMulticastUDP("udp4", nil, &net.UDPAddr{IP: net.IPv4(239, 255, 255, 250), Port: port})
	if err != nil {
		t.Error(err)
	}

	defer con.Close()
	buf := make([]byte, max_buffer)
	n, _, err := con.ReadFromUDPAddrPort(buf)
	if err != nil {
		t.Error(err)
	}

	if string(buf[:n]) != string(b) {
		t.Errorf("upd multicast data did not match expected %s got %s", b, buf[:n])
	}
}
