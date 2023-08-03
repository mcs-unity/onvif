package soap

import "encoding/xml"

type IXmlTag interface {
}

type action struct {
	XMLName xml.Name `xml:"wsa:action"`
	Action  string   `xml:",chardata"`
}

type header struct {
	XMLName xml.Name `xml:"SOAP-ENV:Header"`
	Action  action
	Body    []any
}

type body struct {
	XMLName xml.Name `xml:"SOAP-ENV:Body"`
	Body    []any
}

type Envelop struct {
	XMLName xml.Name `xml:"SOAP-ENV:Envelope"`
	Xmlns   string   `xml:"xmlns:SOAP-ENV,attr"`
	Header  header
	Body    body
}
