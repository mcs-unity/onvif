package digest

import (
	"fmt"
	"testing"
)

func TestDigest(t *testing.T) {
	d, err := Generate("admin", "adminPw")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(d)
}
