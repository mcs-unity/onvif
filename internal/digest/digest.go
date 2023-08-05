package digest

import (
	"github.com/mcs-unity/onvif/internal/soap"
)

func (d Digest) ToXml() ([]byte, error) {
	s := Security{
		Digest: d,
	}
	return soap.WriteXml(s)
}

func (d *Digest) nounce() error {
	s, err := RandomString(12)
	if err != nil {
		return err
	}
	d.Nounce = s
	return err

}
