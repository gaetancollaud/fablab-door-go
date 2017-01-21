package config

import (
	"os"
	"encoding/json"
	"log"
)

const CONFIG_FILE = "config.json"
const USERS_FILE = "users.json"

type User struct {
	Rfid string
	Name string
}

type Configuration struct {
	Serial   string
	Endpoint string
}

func GetConfig() (*Configuration) {
	config := new(Configuration)

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		log.Fatal("Cannot read config file config.json")
	}

	return config
}

func GetUsers() ([]User) {
	var users []User

	if _, err := os.Stat(USERS_FILE); err == nil {
		file, _ := os.Open(USERS_FILE)
		decoder := json.NewDecoder(file)
		err := decoder.Decode(&users)
		if err != nil {
			log.Fatal("Cannot read users file ", USERS_FILE)
		}
	}

	return users;
}
