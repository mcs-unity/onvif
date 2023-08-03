package soap

import "encoding/xml"

type IXmlTag interface {
}

type header struct {
	XMLName xml.Name `xml:"SOAP-ENV:Header"`
}

type body struct {
	XMLName xml.Name `xml:"SOAP-ENV:Body"`
}

type Envelop struct {
	XMLName   xml.Name `xml:"SOAP-ENV:Envelope"`
	Namespace string   `xml:"xmlns:SOAP-ENV,attr"`
	Header    header
	Body      body
}
