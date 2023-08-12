package event

import (
	"fmt"
	"os"
	"testing"
)

func TestEventLogger(t *testing.T) {
	f := GetLogger(EVENT)
	if f != nil {
		t.Error("event log are should never be exported")
	}
}

func TestCreateLogFile(t *testing.T) {
	GetLogger(INFO)
	_, err := os.Open(fmt.Sprintf("./%s.log", INFO))
	if err != nil {
		t.Errorf("%s file was not created\n: %s", INFO, err)
	}

}

func TestNestedEventLogger(t *testing.T) {
	f := GetLogger(ERROR)
	f("I'm an error")
}

func TestCleanup(t *testing.T) {
	files := []string{EVENT, ERROR, INFO}
	for _, v := range files {
		err := os.Remove(fmt.Sprintf("./%s.log", v))
		if err != nil {
			t.Error(err)
		}
	}
}
