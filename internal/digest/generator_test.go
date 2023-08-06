package digest

import (
	"strings"
	"testing"
)

func TestDigest(t *testing.T) {
	_, err := Generate("admin", "adminPw")
	if err != nil {
		t.Error(err)
	}
}

func TestConvertToXML(t *testing.T) {
	d, err := Generate("admin123", "admin123")
	if err != nil {
		t.Error(err)
	}

	xml, err := d.ToXml()
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(xml), "<Security xmlns=") ||
		!strings.Contains(string(xml), "<UsernameToken>") ||
		!strings.Contains(string(xml), "<Username>") ||
		!strings.Contains(string(xml), "<Password Type=") ||
		!strings.Contains(string(xml), "<Nonce EncodingType=") ||
		!strings.Contains(string(xml), "<Created xmlns=") {
		t.Error("Invalid digest payload")
	}
}
