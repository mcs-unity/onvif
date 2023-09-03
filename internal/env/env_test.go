package env

import (
	"os"
	"testing"
)

func TestInvalidFilePath(t *testing.T) {
	if err := LoadVariables(""); err == nil {
		t.Error(err)
	}
}

func TestEmptyWithSpaceFilePath(t *testing.T) {
	if err := LoadVariables(" "); err == nil {
		t.Error(err)
	}
}

func TestLoadEnv(t *testing.T) {
	if err := LoadVariables("./.env"); err != nil {
		t.Error(err)
	}
}

func TestLoadInvalidEnvFile(t *testing.T) {
	if err := LoadVariables("./.env.invalid"); err == nil {
		t.Error(err)
	}
}

func TestGetValueFromKey(t *testing.T) {
	if err := LoadVariables("./.env"); err != nil {
		t.Error(err)
	}

	if v := os.Getenv("hello"); v != "test" {
		t.Error("unable to get key from os")
	}
}
