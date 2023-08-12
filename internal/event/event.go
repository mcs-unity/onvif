package event

import (
	"fmt"
	"log"
	"os"
)

var eventLog *Log
var eventLogs map[EventType]*Log
var wd = ""

func logToEvent(e event) {
	eventLog.log.Println(e)
}

func GetLogger(e EventType) func(e event) {
	if e == EVENT {
		return nil
	}

	var l *Log
	if v, ok := eventLogs[e]; ok {
		l = v
	} else {
		l = add(e)
	}

	return func(e event) {
		logToEvent(e)
		l.log.Println(e)
	}
}

func add(e EventType) *Log {
	f, err := os.OpenFile(fmt.Sprintf("%s/%s.log", wd, e), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	l := Log{
		log: *log.New(f, e+" ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	return &l
}

func init() {
	d, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	wd = d
	eventLogs = make(map[string]*Log)
	eventLog = add(EVENT)
}
