package main

import (
	"./rfid"
	"./logger"
	"./manager"
	"./actuator"
	"os"
	"bufio"
)

func main() {

	logger.ConfigureLogger()


	//if(len(os.Args)<2){
	//	log.Fatal("Usage : fablab-door PORT")
	//}
	//
	//portName := os.Args[1]
	rfIdChannel := make(chan string)
	commandChannel := make(chan actuator.CommandType)

	manager := manager.New(rfIdChannel, commandChannel);
	rfidReader, err := rfid.New(rfIdChannel)
	actuator := actuator.New(commandChannel)

	logInstance := logger.GetLogger("main")

	if err != nil {
		logInstance.Fatal(err)
	}

	defer rfidReader.Stop()

	go rfidReader.Read()
	go manager.Run()
	go actuator.Run()

	logInstance.Printf("All services statted")

	logInstance.Println("Press enter to quit")
	bufio.NewReader(os.Stdin).ReadString('\n')
	logInstance.Println("Exiting")
}



