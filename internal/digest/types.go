package digest

import "encoding/xml"

type IDigest interface {
	ToXml() ([]byte, error)
}

type Security struct {
	Xmlns  string `xml:"attr"`
	Digest Digest
}

type Digest struct {
	XmlName   xml.Name `xml:"UsernameToken"`
	Username  string   `xml:"Username"`
	Password  string   `xml:"Password"`
	Nounce    string   `xml:"Nonce"`
	Timestamp string   `xml:"Created"`
}
