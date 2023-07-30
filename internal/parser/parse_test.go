package parser

import (
	"bytes"
	"fmt"
	"testing"
)

type test struct {
	HI string `xml:"hi,omitempty"`
}

func TestWriteXML(t *testing.T) {
	data := test{
		HI: "Hello",
	}

	b, err := WriteXml(data)
	if err != nil {
		t.Error(data)
	}

	fmt.Printf("%s", b)

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
