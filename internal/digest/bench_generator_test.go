package digest

import "testing"

func BenchmarkDigest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate("admin", "adminPw")
	}
}

func BenchmarkConvertToXML(b *testing.B) {
	d, err := Generate("admin123", "admin123")
	if err != nil {
		b.Error(err)
	}
	for i := 0; i < b.N; i++ {
		_, err = d.ToXml()
		if err != nil {
			b.Error("test failed")
		}
	}
}
