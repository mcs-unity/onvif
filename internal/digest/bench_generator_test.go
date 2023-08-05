package digest

import "testing"

func BenchmarkDigest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate("admin", "adminPw")
	}
}
