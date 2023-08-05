package digest

import "encoding/xml"

type IDigest interface {
	ToXml() ([]byte, error)
	nounce() error
	hash()
}

type Security struct {
	Xmlns  string `xml:"xmlns,attr"`
	Digest Digest
}

type Digest struct {
	XMLName   xml.Name `xml:"UsernameToken"`
	Username  string   `xml:"Username"`
	Password  string   `xml:"Password"`
	Nounce    string   `xml:"Nonce"`
	Timestamp string   `xml:"Created"`
}
