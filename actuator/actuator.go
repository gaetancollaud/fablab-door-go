package actuator

import (
	"../logger"
	"../config"
	"github.com/Sirupsen/logrus"
	"net"
	"time"
	"fmt"
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
	connection     net.Conn
}

func New(channel chan CommandType) (*Actuator) {
	actuator := new(Actuator)
	actuator.log = logger.GetLogger("Acutator")
	actuator.commandChannel = channel

	actuator.log.Level = logrus.DebugLevel

	return actuator
}

func (actuator *Actuator)Run() {
	actuator.openConnection()
	defer actuator.connection.Close();
	for command := range actuator.commandChannel {
		actuator.handleCommand(command)
	}
}

func (actuator *Actuator)handleCommand(command CommandType) {
	actuator.log.Debug("Command received ", command)

	var msg string
	msg = "01";

	fmt.Println(actuator.connection, msg);

}

func (actuator *Actuator)openConnection() {

	var err error
	connected := false
	for !connected {

		host := config.GetConfig().Ipx800Endpoint.Url
		actuator.log.Info("Try conecting to ipx at ", host)
		actuator.connection, err = net.Dial("tcp", host)
		if (err != nil) {
			actuator.log.Error("Cannot connect to IPX800 at ", host, ". Retry in 10sec")
			time.Sleep(10 * time.Second)
		} else {
			actuator.log.Info("Connexion successfull to IPX at ", host)
			connected = true
		}
	}

}