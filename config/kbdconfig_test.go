package config

import (
	"kbd/module"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setUP()
	m.Run()
	tearDown()
}

func setUP() {
	os.Setenv("KBD", "../testdata")
}

func tearDown() {
	os.Unsetenv("KBD")
	kubeConfig := module.KbdConfig{
		AppName:    "kbd",
		AppVersion: "0.0.1",
		Realms:     []module.Realm{{Name: "kind-docker-desktop", Desc: "kind k8s cluster", KubeConfig: "config", Active: false}},
	}
	if err := SyncToConfigFile(kubeConfig); err != nil {
		log.Printf("teardown failed: %v", err)
	}
}

func TestLoadFromConfigFile(t *testing.T) {
	want := module.KbdConfig{
		AppName:    "kbd",
		AppVersion: "0.0.1",
		Realms:     []module.Realm{{Name: "kind-docker-desktop", Desc: "kind k8s cluster", KubeConfig: "config", Active: false}},
	}
	kbdConfig, err := LoadFromConfigFile()

	if err != nil {
		t.FailNow()
	}
	if kbdConfig.AppName != want.AppName {
		t.Errorf("kdb config %s != %s", want.AppName, kbdConfig.AppName)
	}
	if kbdConfig.AppVersion != want.AppVersion {
		t.Errorf("kdb config %s != %s", want.AppVersion, kbdConfig.AppVersion)
	}
	if len(kbdConfig.Realms) == 0 || kbdConfig.Realms[0] != want.Realms[0] {
		t.Errorf("kdb config %v != %v", want.Realms[0], kbdConfig.Realms[0])
	}
}

func TestSyncToConfigFile(t *testing.T) {
	kConfig, err := LoadFromConfigFile()
	if err != nil {
		t.FailNow()
	}

	kConfig.Realms = append(kConfig.Realms, module.Realm{Name: "test", Desc: "desc", KubeConfig: "test_config"})
	err = SyncToConfigFile(kConfig)
	if err != nil {
		t.FailNow()
	}

	kConfig, err = LoadFromConfigFile()
	if err != nil {
		t.FailNow()
	}
	if kConfig.Realms[1] != kConfig.Realms[1] {
		t.Errorf("kdb config %v != %v", kConfig.Realms[1], kConfig.Realms[1])
	}
}
