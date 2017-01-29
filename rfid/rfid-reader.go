package rfid

import (
	"github.com/tarm/serial"
	"bytes"
	"strings"
	"../config"
	"../logger"
	"github.com/Sirupsen/logrus"
)

type RfidReader struct {
	config   *serial.Config
	openPort *serial.Port
	log *logrus.Entry
	rfidChannel chan string
}

func New(rfidChannel chan string) (*RfidReader, error) {
	var err error
	serialPort := config.GetConfig().Serial

	rfid := new(RfidReader)
	rfid.log = logger.GetLogger("RfidReader")
	rfid.rfidChannel = rfidChannel
	rfid.config = &serial.Config{Name: serialPort, Baud: 9600}
	rfid.openPort, err = serial.OpenPort(rfid.config)
	rfid.log.Printf("Port %s open", serialPort)
	if err != nil {
		return nil, err
	}
	return rfid, nil
}

func (rfid *RfidReader)Read() {
	var buffer bytes.Buffer
	buf := make([]byte, 128)
	for true {
		n, err := rfid.openPort.Read(buf)
		if err != nil {
			rfid.log.Fatal(err)
		}
		buffer.Write(buf[:n])

		buffer = rfid.analyseString(buffer);
	}
}

func (rfid *RfidReader)Stop() {
	rfid.openPort.Close()
}

func (rfid *RfidReader)analyseString(buffer bytes.Buffer) (bytes.Buffer) {
	str := buffer.String()
	index := strings.Index(str, "\r\n")
	if (index != -1) {
		id := str[:index]
		if (len(id) > 0) {
			rfid.handleId(id);
		}
		if (len(str) == index + 2) {
			str = ""
		} else {
			str = str[index + 2:]
		}
	}
	var ret bytes.Buffer;
	ret.WriteString(str);
	return ret;
}

func (rfid *RfidReader) handleId(id string) {
	rfid.log.Debug("Read TAG ID: %s", id)
	rfid.rfidChannel <- id
}
