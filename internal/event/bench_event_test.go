package event

import (
	"fmt"
	"os"
	"testing"
)

func BenchmarkEventLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f := GetLogger(EVENT)
		if f != nil {
			b.Error("event log are should never be exported")
		}
	}
}

func BenchmarkCreateLogFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetLogger(INFO)
	}
}

func BenchmarkNestedEventLogger(b *testing.B) {
	f := GetLogger(ERROR)
	for i := 0; i < b.N; i++ {
		f("I'm an error")
	}
}

func BenchmarkCleanup(b *testing.B) {
	files := []string{EVENT, ERROR, INFO}
	for _, v := range files {
		err := os.Remove(fmt.Sprintf("./%s.log", v))
		if err != nil {
			b.Error(err)
		}
	}
}
