package manager

import (
	"../config"
	"strings"
	"gopkg.in/resty.v0"
	"../logger"
	"../actuator"
	"github.com/Sirupsen/logrus"
)

type Manager struct {
	log *logrus.Entry
	rfidChannel chan string
	commandChannel chan actuator.CommandType
}

func New(rfidChannel chan string, commandChannel chan actuator.CommandType) (*Manager) {
	manager := new(Manager)
	manager.log = logger.GetLogger("Manager")
	manager.rfidChannel = rfidChannel
	manager.commandChannel = commandChannel
	return manager
}

func (manager *Manager)Run(){
	for rfid := range manager.rfidChannel {
		manager.checkRfid(rfid);
	}
}

func (manager *Manager)checkRfid(rfid string) {
	users := config.GetUsers()

	found := false

	for _, user := range users {
		if strings.EqualFold(rfid, user.Rfid) {
			found = true
		}
	}

	if (found) {
		manager.log.Infof("RFID %v found in the users list", rfid)
		manager.rfidAccepted(rfid);
	} else {
		manager.log.Infof("RFID %v not found in the users list, trying to call the server", rfid)
		//maybe this user was just added, check online
		if (manager.checkRfidOnline(rfid)) {
			manager.log.Infof("Server said that the RFID %v is allowed", rfid)
			manager.rfidAccepted(rfid);
		} else {
			manager.log.Infof("Server doesn't know about the RFID %v either", rfid)
			manager.rfidRejected(rfid);
		}
	}
}

func (manager *Manager)rfidAccepted(rfid string){
	manager.commandChannel <- actuator.COMMAND_OPEN_SHORT
	//send event
}

func (manager *Manager)rfidRejected(rfid string){
	//send event
}

func (manager *Manager)checkRfidOnline(rfid string) (bool) {
	endpoint := config.GetConfig().FablabEndpoint
	resty.SetBasicAuth(endpoint.User, endpoint.Password)
	resp, err := resty.R().Get(endpoint.Url + "/api/v1/system/door/allowed?rfid=" + rfid)

	if (err != nil) {
		manager.log.Errorf("Error while checking rfid " + rfid, err)
		return false
	}
	if (err == nil && resp.StatusCode() == 200) {
		return strings.EqualFold("true", resp.String())
	}

	manager.log.Errorf("Error while checking for rfid '%v', response status=%v, body=%v", rfid, resp.StatusCode(), resp)
	return false
}
