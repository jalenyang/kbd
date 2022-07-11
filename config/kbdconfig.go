package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	rlm "kbd/realm"
	"log"
	"os"
)

var kbdConfig KbdConfig

type KbdConfig struct {
	AppName    string      `yaml:"appName"`
	AppVersion string      `yaml:"appVersion"`
	Realms     []rlm.Realm `yaml:"realms"`
}

// LoadFromConfigFile parseConfig load the app config
func LoadFromConfigFile(configFile string) (KbdConfig, error) {
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

func SyncToConfigFile(configFile string) error {
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
