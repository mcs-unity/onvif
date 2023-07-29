package connection

import (
	"bytes"
	"fmt"
	"net/http"
)

func Post(url string, t ContentType, payload []byte) error {
	c, err := http.Post(url, t, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	if c.StatusCode != http.StatusOK {
		return fmt.Errorf("request returned status code %d", c.StatusCode)
	}

	return nil
}
