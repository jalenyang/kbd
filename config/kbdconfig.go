package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"kbd/module"
	"log"
	"os"
	"path"
)

func getConfigFile() string {
	configPath := os.Getenv("KBD")
	if configPath != "" {
		return path.Join(configPath, "kbd.yaml")
	}
	home, err := os.UserHomeDir()
	if err != nil {
		log.Printf("failed to get the home dir %v", err)
	}
	return path.Join(home, "kbd.yaml")
}

// LoadFromConfigFile parseConfig load the app config
func LoadFromConfigFile() (module.KbdConfig, error) {
	var kbdConfig module.KbdConfig
	configFile := getConfigFile()
	fi, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Panicln("Error: failed to open the config file 'kdb.yaml'")
	}
	defer fi.Close()

	data, err := ioutil.ReadAll(fi)
	if err != nil {
		log.Panicln("Error: failed to read the config file 'kdb.yaml'")
	}
	err = yaml.Unmarshal(data, &kbdConfig)

	if err != nil {
		log.Printf("Error: failed to unmarshal the config file 'kdb.yaml' for %v", err)
	}
	return kbdConfig, err
}

func SyncToConfigFile(kbdConfig module.KbdConfig) error {
	configFile := getConfigFile()
	fi, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Panicln("Error: failed to open the config file 'kdb.yaml'")
	}
	defer fi.Close()

	out, err := yaml.Marshal(kbdConfig)
	if err != nil {
		log.Panicln("Error: failed to sync the configurations from the memory to the config file 'kdb.yaml'")
		return err
	}
	if _, err := fi.Write(out); err != nil {
		return err
	}
	return nil
}
