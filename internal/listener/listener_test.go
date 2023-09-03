package listener

import "testing"

var invalidIP = ""
var validIP = "127.0.0.1"
var port1 uint16 = 8080
var port2 uint16 = 8090

func TestInvalidListen(t *testing.T) {
	_, err := Listener(invalidIP, port1, TCP)
	if err == nil {
		t.Error("did not catch invalid ip address")
	}
}

func TestCloseListener(t *testing.T) {
	_, err := Listener(validIP, port1, TCP)
	if err != nil {
		t.Error(err)
	}

	if l := Get(port1); l == nil {
		t.Error("unable to find listener listening on port:", port1)
	}

	if err := Close(port1); err != nil {
		t.Error(err)
	}
}

func TestTwoListenerOnSamePort(t *testing.T) {
	_, err := Listener(validIP, port1, TCP)
	defer Close(port1)
	if err != nil {
		t.Error(err)
	}

	_, err = Listener(validIP, port1, TCP)
	if err == nil {
		t.Error("a second listener was opened with the same port 8080")
	}
}

func TestOpenTwoListener(t *testing.T) {
	_, err := Listener(validIP, port1, TCP)
	defer Close(port1)
	if err != nil {
		t.Error(err)
	}

	_, err = Listener(validIP, port2, TCP)
	defer Close(port2)
	if err != nil {
		t.Error(err)
	}

	connections := GetList(port1, port2)
	if len(connections) < 2 {
		t.Error("list should contain two listeners")
	}
}

func TestGetPort(t *testing.T) {
	l, err := Listener(validIP, port1, TCP)
	if err != nil {
		t.Error(err)
	}
	defer Close(port1)

	port := GetPort(l)
	if port != port1 {
		t.Errorf("%d did not match %d", port, port1)
	}
}
