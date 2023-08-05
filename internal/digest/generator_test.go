package digest

import (
	"fmt"
	"testing"
)

func TestDigest(t *testing.T) {
	d, err := Generate("admin", "adminPw")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(d)
}

func TestConvertToXML(t *testing.T) {
	d, err := Generate("admin", "adminPw")
	if err != nil {
		t.Error(err)
	}

	xml, err := d.ToXml()
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(xml))
}
