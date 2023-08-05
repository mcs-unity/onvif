package digest

import (
	"crypto/sha1"
	"encoding/base64"

	"github.com/mcs-unity/onvif/internal/soap"
)

func (d Digest) ToXml() ([]byte, error) {
	s := Security{
		Xmlns:  "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd",
		Digest: d,
	}
	return soap.WriteXml(s)
}

func (d *Digest) hash() {
	sum := sha1.Sum([]byte(d.Nounce + d.Timestamp + d.Password))
	d.Password = base64.RawStdEncoding.EncodeToString(sum[0:])
	d.Nounce = base64.RawStdEncoding.EncodeToString([]byte(d.Nounce))
}

func (d *Digest) nounce() error {
	s, err := RandomString(12)
	if err != nil {
		return err
	}
	d.Nounce = s
	return err

}
