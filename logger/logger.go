package logger

import (
	"github.com/Sirupsen/logrus"
	"os"
	"io"
	"log"
)


//var logInstance *log.Logger
var writter io.Writer

func ConfigureLogger() {
	file, err := os.OpenFile("fablab-door.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", ":", err)
	}

	writter = io.MultiWriter(os.Stdout, file)
}

func GetLogger(name string) (*logrus.Entry) {
	log := logrus.WithFields(logrus.Fields{
		"package": name,
	})
	log.Logger.Out = writter
	return log;
}