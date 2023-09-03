package env

import (
	"bytes"
	"errors"
	"os"
	"strings"
)

func splitVariables(b []byte) error {
	n := bytes.Split(b, []byte("\n"))
	for _, v := range n {
		if len(v) == 0 {
			continue
		}

		if kv := bytes.SplitN(v, []byte("="), 2); len(kv) == 2 {
			if err := os.Setenv(string(kv[0]), string(kv[1])); err != nil {
				return err
			}
		} else {
			return errors.New("invalid key has been parsed")
		}
	}
	return nil
}

func LoadVariables(path string) error {
	if t := strings.Trim(path, ""); t == "" {
		return errors.New("empty file path")
	}

	f, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = splitVariables(f)
	if err != nil {
		return err
	}

	return nil
}
