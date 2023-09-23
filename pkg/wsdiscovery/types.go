package wsdiscovery

import (
	"time"
)

const max_delay = 500 * time.Millisecond
const max_wait = 5 * time.Second
const max_matches = 100

const multicastIP = "239.255.255.250"
const port = 3702
const max_buffer = 8192

type Match struct {
	IP   string `json:"ip,omitempty"`
	Data []byte `json:"data,omitempty"`
}
