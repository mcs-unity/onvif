package digest

import (
	"crypto/rand"
	"encoding/base32"
	"time"
)

func RandomString(len uint8) (string, error) {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base32.HexEncoding.EncodeToString(b), nil
}

func Generate(u, p string) (IDigest, error) {
	d := Digest{
		Username:  u,
		Password:  p,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	err := d.nounce()
	if err != nil {
		return nil, err
	}

	d.hash()
	return &d, nil
}
