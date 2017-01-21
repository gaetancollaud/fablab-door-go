package main

import (
	"log"
	"./rfid"
	"time"
)

func main() {
	reader, err := rfid.Start("COM4")
	if err != nil {
		log.Fatal(err)
	}

	defer reader.Stop()

	go reader.Read()

	log.Printf("Started main")

	log.Println("Waiting for input")
	for true {
		time.Sleep(1000 * time.Millisecond)
	}
	log.Println("Exiting")
}


