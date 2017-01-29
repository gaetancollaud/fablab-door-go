package actuator

import (
	"../logger"
	"github.com/Sirupsen/logrus"
)

type CommandType int

const (
	COMMAND_OPEN_SHORT CommandType = iota
	COMMAND_OPEN
	COMMAND_CLOSE
	COMMAND_EXIT
)

type Actuator struct {
	log            *logrus.Entry
	commandChannel chan CommandType
}

func New(channel chan CommandType) (*Actuator) {
	actuator := new(Actuator)
	actuator.log = logger.GetLogger("Acutator")
	actuator.commandChannel = channel
	return actuator
}

func (actuator *Actuator)Run() {
	for command := range actuator.commandChannel {
		actuator.handleCommand(command)
	}
}

func (actuator *Actuator)handleCommand(command CommandType) {
	actuator.log.Debug("Command received ", command)
}