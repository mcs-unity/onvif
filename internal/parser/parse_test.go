package parser

import (
	"bytes"
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

	if !bytes.Contains(b, []byte("</hi>")) {
		t.Error("xml was unable to pass struct")
	}
}

func BenchmarkWriteXML(b *testing.B) {
	data := test{
		HI: "Hello",
	}

	for i := 0; i < b.N; i++ {
		WriteXml(data)
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

func BenchmarkParseXML(b *testing.B) {
	xml := "<test><hi>Hello</hi></test>"
	data := test{}
	for i := 0; i < b.N; i++ {
		ParseXml([]byte(xml), &data)
	}
}
