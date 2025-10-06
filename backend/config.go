package backend

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Config() {
	getConfig()
	readConfig()
}

func readConfig() (string, string) {
	f := getConfig()
	file, err := os.Open(f)
	if err != nil {
		log.Fatalln("Error getting file ", err)
	}
	defer file.Close()

	bValue, _ := io.ReadAll(file)
	var c ConfigStruct

	json.Unmarshal(bValue, &c)
	return c.Ip, c.Port
}

func getConfig() string {
	configPath, err := os.UserConfigDir()
	if err != nil {
		log.Fatalln("Error getting config dir ", err)
	}

	configDir := filepath.Join(configPath, configFolder)
	if _, err = os.Stat(configDir); os.IsNotExist(err) {
		log.Print("Couldn't get config folder creating for you ", configDir)
		err = os.Mkdir(configDir, 0755)
		if err != nil {
			log.Fatalln("Error creating config file")
		}
	}

	configFile := filepath.Join(configDir, configName)
	if _, err = os.Stat(configFile); os.IsNotExist(err) {
		log.Print("Couldn't get config folder creating for you ", configFile)
		_, err = os.Create(configFile)
		if err != nil {
			log.Fatalln("Error creating config file")
		}
		writeConfig()
	}

	return configFile
}

func writeConfig() {
	defaultVal := &ConfigStruct{
		Ip: "0.0.0.0",
		Port: "31311",
	}
	bytes, _ := json.MarshalIndent(defaultVal, "", "\t")
	os.WriteFile(getConfig(), bytes, 0644)
}