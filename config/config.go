package config

import (
	"os"
	"encoding/json"
	"../logger"
)

var configLogger = logger.GetLogger("Config");

const CONFIG_FILE = "config.json"
const USERS_FILE = "users.json"

type User struct {
	Rfid string
	Name string
}

type Configuration struct {
	Serial   string
	FablabEndpoint ConfigurationEndpoint
}

type ConfigurationEndpoint struct {
	Url string
	User string
	Password string
}

func GetConfig() (*Configuration) {
	config := new(Configuration)

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)
	if err != nil {
		configLogger.Fatal("Cannot read config file config.json")
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
			configLogger.Fatal("Cannot read users file ", USERS_FILE)
		}
	}

	return users;
}
