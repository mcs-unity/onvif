package digest

import "encoding/xml"

type IDigest interface {
	ToXml() ([]byte, error)
}

type Security struct {
	Xmlns          string `xml:"xmlns,attr"`
	MustUnderstand uint8  `xml:"s:mustUnderstand,attr"`
	Digest         Digest
}

type password struct {
	Type     string `xml:"Type,attr"`
	Password string `xml:",chardata"`
}

type nonce struct {
	EncodingType string `xml:"EncodingType,attr"`
	Nonce        string `xml:",chardata"`
}

type timestamp struct {
	Xmlns   string `xml:"xmlns,attr"`
	Created string `xml:",chardata"`
}

type Digest struct {
	XMLName  xml.Name `xml:"UsernameToken"`
	Username string   `xml:"Username"`
	Password password
	Nonce    nonce
	Created  timestamp
}
