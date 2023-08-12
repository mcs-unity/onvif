package event

import "log"

type EventType = string
type event = string

const (
	ERROR EventType = "error"
	INFO  EventType = "info"
	EVENT EventType = "event"
)

type Log struct {
	log log.Logger
}
