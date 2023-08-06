package digest

import (
	"crypto/sha1"
	"encoding/base64"

	"github.com/mcs-unity/onvif/internal/soap"
)

func (d Digest) ToXml() ([]byte, error) {
	s := Security{
		Xmlns:          "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
		Digest:         d,
		MustUnderstand: 1,
	}
	return soap.WriteXml(s)
}

func (d *Digest) hash() error {
	s, err := RandomBytes(12)
	if err != nil {
		return err
	}
	sum := sha1.Sum([]byte(string(s) + d.Created.Created + d.Password.Password))
	d.Password.Password = base64.RawStdEncoding.EncodeToString(sum[0:])
	d.Nonce.Nonce = base64.RawStdEncoding.EncodeToString(s)
	return nil
}
