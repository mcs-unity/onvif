package wsdiscovery

import (
	"fmt"
	"net"
	"time"
)

// http://schemas.xmlsoap.org/ws/2005/04/discovery namespace must be included in the soap request
// send multicast message within a network
// PORT 3702, IPV4 239.255.255.250, IPV6 FF02::C
// Message is sent over UDP Following SOAP over UDP spec
// Target service will send a message with a delay that is random between 0,500 ms
// Max timeout for a response is set to 5 seconds once a multicast is broadcasted in the network

func SendMulticastProbe(data []byte) ([]Match, error) {
	con, err := net.ListenPacket("udp4", "0.0.0.0:1024")
	if err != nil {
		return nil, err
	}
	defer con.Close()

	dst, err := net.ResolveUDPAddr("udp4", fmt.Sprintf("%s:%d", multicastIP, port))
	if err != nil {
		return nil, err
	}

	if err := con.SetWriteDeadline(time.Now().Add(max_delay)); err != nil {
		return nil, err
	}

	_, err = con.WriteTo(data, dst)
	if err != nil {
		return nil, err
	}

	b := make([]byte, max_buffer)
	m := make([]Match, 0, max_matches)

	if err := con.SetReadDeadline(time.Now().Add(max_wait)); err != nil {
		return nil, err
	}

	for {
		n, cm, err := con.ReadFrom(b)
		if err != nil {
			return m, err
		}
		d := Match{
			IP:   cm.String(),
			Data: b[:n],
		}
		m = append(m, d)
	}
}
