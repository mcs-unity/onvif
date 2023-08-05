package digest

import (
	"crypto/rand"
	"time"
)

func RandomBytes(len uint8) ([]byte, error) {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func Generate(u, p string) (IDigest, error) {
	d := Digest{
		Username:  u,
		Password:  p,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	err := d.hash()
	if err != nil {
		return nil, err
	}

	return &d, nil
}
