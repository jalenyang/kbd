package module

type KbdConfig struct {
	AppName    string  `yaml:"appName"`
	AppVersion string  `yaml:"appVersion"`
	Realms     []Realm `yaml:"realms"`
}

type Realm struct {
	Name       string `yaml:"name"`
	Desc       string `yaml:"desc"`
	KubeConfig string `yaml:"kubeConfig"`
	Active     bool   `yaml:"active"`
	E2eTestApi string `yaml:"e2eTestApi"`
}
