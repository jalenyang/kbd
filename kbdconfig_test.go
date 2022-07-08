package main

import (
	"testing"
)

func TestLoadFromConfigFile(t *testing.T) {
	want := KbdConfig{
		AppName:    "kbd",
		AppVersion: "0.0.1",
		Realms:     []Realm{{Name: "kind-docker-desktop", Desc: "kind k8s cluster", KubeConfig: "config"}},
	}
	kbdConfig, err := loadFromConfigFile("testdata/kbd.yaml")

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
		t.Errorf("kdb config %s != %s", want.Realms[0], kbdConfig.Realms[0])
	}
}

func TestSyncToConfigFile(t *testing.T) {
	kConfig, err := loadFromConfigFile("testdata/kbd.yaml")
	if err != nil {
		t.FailNow()
	}

	kbdConfig.Realms = append(kConfig.Realms, Realm{Name: "test", Desc: "desc", KubeConfig: "test_config"})
	err = syncToConfigFile("testdata/kbd.yaml")
	if err != nil {
		t.FailNow()
	}

	kConfig, err = loadFromConfigFile("testdata/kbd.yaml")
	if err != nil {
		t.FailNow()
	}
	if kConfig.Realms[1] != kbdConfig.Realms[1] {
		t.Errorf("kdb config %v != %v", kConfig.Realms[1], kbdConfig.Realms[1])
	}
}
