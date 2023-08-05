package digest

import (
	"time"

	"github.com/mcs-unity/onvif/internal/soap"
)

func (d Digest) ToXml() ([]byte, error) {
	s := Security{
		Digest: d,
	}
	return soap.WriteXml(s)
}

func Generate(u, p string) (IDigest, error) {
	d := Digest{
		Username:  u,
		Password:  p,
		Nounce:    "",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	return d, nil
}
