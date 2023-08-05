package digest

import (
	"testing"
)

func TestDigest(t *testing.T) {
	_, err := Generate("admin", "adminPw")
	if err != nil {
		t.Error(err)
	}
}
