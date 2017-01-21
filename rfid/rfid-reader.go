package rfid

import (
	"log"
	"github.com/tarm/serial"
	"bytes"
	"strings"
)

type Rfid struct {
	config   *serial.Config
	openPort *serial.Port
}

func Start(port string) (*Rfid, error) {
	var err error
	rfid := new(Rfid)
	rfid.config = &serial.Config{Name: port, Baud: 9600}
	rfid.openPort, err = serial.OpenPort(rfid.config)
	log.Printf("Port %s open", port)
	if err != nil {
		return nil, err
	}
	return rfid, nil
}

func (rfid *Rfid)Read() {
	var buffer bytes.Buffer
	buf := make([]byte, 128)
	for true {
		n, err := rfid.openPort.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		buffer.Write(buf[:n])

		buffer = AnalyseString(buffer);
	}
}

func (rfid *Rfid)Stop() {
	rfid.openPort.Close()
}

func AnalyseString(buffer bytes.Buffer) (bytes.Buffer) {
	str := buffer.String()
	index := strings.Index(str, "\r\n")
	if (index != -1) {
		id := str[:index]
		if (len(id) > 0) {
			HandleId(id);
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

func HandleId(id string) {
	log.Printf("Read TAG ID: %s", id)
}
