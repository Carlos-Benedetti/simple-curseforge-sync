package config

import (
	"log"

	"gopkg.in/ini.v1"
)

var FILE_NAME = "config.ini"
var LAST_INSTANCE_KEY = "last_instance"

func getIni() *ini.File {
	cfg, err := ini.Load(FILE_NAME)
	if err != nil {
		cfg = createIni()
	}
	return cfg

}
func createIni() *ini.File {
	cfg := ini.Empty()
	if err := cfg.SaveTo(FILE_NAME); err != nil {
		log.Fatalf("Failed to save INI file: %v", err)
	}

	log.Println("config.ini file created successfully.")
	return cfg
}

func saveIni(cfg *ini.File) {
	if err := cfg.SaveTo(FILE_NAME); err != nil {
		log.Fatalf("Failed to save INI file: %v", err)
	}

	log.Println("config.ini updated successfully.")
}

func GetLastSession() *string {
	var value string
	cfg := getIni()
	rootSession := cfg.Section("")
	hasLastInstance := rootSession.HasKey(LAST_INSTANCE_KEY)

	if hasLastInstance {
		value = rootSession.Key(LAST_INSTANCE_KEY).Value()
		log.Println("value: ", *&value)
	}

	return &value
}

func SetLastCSession(value string) {
	cfg := getIni()
	cfg.Section("").Key(LAST_INSTANCE_KEY).SetValue(value)
	saveIni(cfg)
}
