package main

import "log"

type Realm struct {
	Name       string `yaml:"name"`
	Desc       string `yaml:"desc"`
	KubeConfig string `yaml:"kubeConfig"`
}

func getAllRealms() ([]Realm, error) {
	return nil, nil
}

func createRealm(realm Realm) error {
	if realm.Name == "" || realm.KubeConfig == "" {
		log.Println("Invalid realm, realm name and kubeconfig can't be empty")
	}
	//realms := append(kbdConfig.realms, realm)
	//kbdConfig.realms = realms
	return nil
}
