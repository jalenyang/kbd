package main

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

type KbdConfig struct {
	AppName    string  `yaml:"appName"`
	AppVersion string  `yaml:"appVersion"`
	Realms     []Realm `yaml:"realms"`
}

// parseConfig load the app config
func loadFromConfigFile(configFile string) (KbdConfig, error) {
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

func syncToConfigFile(configFile string) error {
	fi, err := os.OpenFile(configFile, os.O_RDWR|os.O_CREATE, 0755)
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
