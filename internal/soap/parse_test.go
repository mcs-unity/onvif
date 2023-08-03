package soap

import (
	"bytes"
	"encoding/xml"
	"testing"
)

type test struct {
	XMLName xml.Name `xml:"plant"`
	HI      string   `xml:"hi,omitempty"`
	Env     string   `xml:"env,attr"`
}

func TestWriteXML(t *testing.T) {
	data := test{
		HI:  "Hello",
		Env: "test",
	}

	b, err := writeXml(data)
	if err != nil {
		t.Error(data)
	}

	if !bytes.Contains(b, []byte("</hi>")) {
		t.Error("xml was unable to pass struct")
	}
}

func TestParseXML(t *testing.T) {
	xml := "<test><hi>Hello</hi></test>"
	data := test{}
	if err := ParseXml([]byte(xml), &data); err != nil {
		t.Error(err)
	}

	if data.HI != "Hello" {
		t.Error("failed to parse XML to struct")
	}
}

func TestSoapTags(t *testing.T) {
	soap := SoapBody()
	b, err := writeXml(soap)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Contains(b, []byte("SOAP-ENV:Envelope xmlns:SOAP-ENV=\"")) ||
		!bytes.Contains(b, []byte("<SOAP-ENV:Header>")) ||
		!bytes.Contains(b, []byte("<SOAP-ENV:Body>")) {
		t.Error("SOAP body is invalid was unable to pass struct")
	}
}
