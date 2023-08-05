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
		Username: u,
		Password: password{
			Type: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest",
		},
		Created: timestamp{
			Xmlns:   "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd",
			Created: time.Now().UTC().Format(time.RFC3339),
		},
		Nonce: nonce{
			EncodingType: "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary",
		},
	}

	err := d.hash()
	if err != nil {
		return nil, err
	}

	return &d, nil
}
