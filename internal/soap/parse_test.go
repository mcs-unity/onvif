package soap

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"strings"
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

	b, err := WriteXml(data)
	if err != nil {
		t.Error(data)
	}

	if !bytes.Contains(b, []byte("</hi>")) {
		t.Error("xml was unable to pass struct")
	}
}

func TestParseXML(t *testing.T) {
	xml := "<plant><hi>Hello</hi></plant>"
	data := test{}
	if err := ParseXml([]byte(xml), &data); err != nil {
		t.Error(err)
	}

	if data.HI != "Hello" {
		t.Error("failed to parse XML to struct")
	}
}

func TestSoapTags(t *testing.T) {
	soap := New("sampleAction")
	b, err := soap.ToString()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(b))

	if !bytes.Contains(b, []byte("SOAP-ENV:Envelope xmlns:SOAP-ENV=\"")) ||
		!bytes.Contains(b, []byte("<SOAP-ENV:Header>")) ||
		!bytes.Contains(b, []byte("<SOAP-ENV:Body>")) {
		t.Error("SOAP body is invalid was unable to pass struct")
	}

	if !bytes.Contains(b, []byte("<wsa:action>sampleAction</wsa:action>")) {
		t.Error("SOAP body is invalid unable to find soap action")
	}
}

func TestAppendHeader(t *testing.T) {
	te := test{}
	te2 := test{}
	s := New("testAction")
	s.AppendToHeader(te, te2)
	b, err := s.ToString()
	fmt.Println(string(b))
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(b), "") {
		t.Error("failed to append header value")
	}
}

func TestAppendBody(t *testing.T) {
	te := test{}
	s := New("testAction")
	s.AppendToBody(te)
	b, err := s.ToString()
	if err != nil {
		t.Error(err)
	}

	if !strings.Contains(string(b), "") {
		t.Error("failed to append bhody value")
	}
}
