package event

import (
	"../logger"
)
var log = logger.GetLogger("EventLogger");

type RfidEventAction string

const (
	EVENT_OPEN RfidEventAction = "OPEN"
	EVENT_CLOSE RfidEventAction = "CLOSE"
	EVENT_TRY_OPEN_BUT_FAIL RfidEventAction = "TRY_OPEN_BUT_FAIL",
)

type RfidEvent struct {
	rfid   string
	action RfidEventAction
}

type EventLogger struct {
	eventChannel chan RfidEvent
}

func New() (*EventLogger) {
	el := new(EventLogger)
	return el;
}

func (el *EventLogger)Run() {
	for event := range el.eventChannel {
		el.handleEvent(event)
	}
}

func (el*EventLogger)handleEvent(event RfidEvent) {

}



