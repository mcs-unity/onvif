package soap

import (
	"testing"
)

func BenchmarkWriteXML(b *testing.B) {
	data := test{
		HI: "Hello",
	}

	for i := 0; i < b.N; i++ {
		writeXml(data)
	}
}

func BenchmarkParseXML(b *testing.B) {
	xml := "<test><hi>Hello</hi></test>"
	data := test{}
	for i := 0; i < b.N; i++ {
		ParseXml([]byte(xml), &data)
	}
}

func BenchmarkAppendHeader(b *testing.B) {
	te := test{}
	te2 := test{}
	s := New("testAction")
	for i := 0; i < b.N; i++ {
		s.AppendToHeader(te, te2)
	}
}

func BenchmarkAppendBody(b *testing.B) {
	te := test{}
	s := New("testAction")
	for i := 0; i < b.N; i++ {
		s.AppendToBody(te)
	}
}
