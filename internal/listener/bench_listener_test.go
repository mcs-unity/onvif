package listener

import (
	"testing"
)

func BenchmarkOpenCloseListener(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := Listen(validIP, port1, TCP)
		if err != nil {
			b.Error(err)
		}

		if err := Close(port1); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGetListener(b *testing.B) {
	_, err := Listen(validIP, port1, TCP)
	if err != nil {
		b.Error(err)
	}
	defer Close(port1)

	for i := 0; i < b.N; i++ {
		Get(port1)
	}
}

func BenchmarkGetListeners(b *testing.B) {
	_, err := Listen(validIP, port1, TCP)
	defer Close(port1)
	if err != nil {
		b.Error(err)
	}

	_, err = Listen(validIP, port2, TCP)
	defer Close(port2)
	if err != nil {
		b.Error(err)
	}

	for i := 0; i < b.N; i++ {
		GetList(port1, port2)
	}
}

func BenchmarkGetPort(b *testing.B) {
	l, err := Listen(validIP, port1, TCP)
	if err != nil {
		b.Error(err)
	}
	defer Close(port1)

	for i := 0; i < b.N; i++ {
		GetPort(l)
	}
}
