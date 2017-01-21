package main

import (
	"log"
	"./rfid"
	"os"
	"bufio"
)

func main() {
	if(len(os.Args)<2){
		log.Fatal("Usage : fablab-door PORT")
	}

	portName := os.Args[1]

	rfidReader, err := rfid.Start(portName)
	if err != nil {
		log.Fatal(err)
	}

	defer rfidReader.Stop()

	go rfidReader.Read()

	log.Printf("All services statted")

	log.Println("Press enter to quit")
	bufio.NewReader(os.Stdin).ReadString('\n')
	log.Println("Exiting")
}


