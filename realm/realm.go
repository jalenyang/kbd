package realm

import (
	"flag"
	"fmt"
	"log"
)

var (
	ram Realm
)

type Realm struct {
	Name       string `yaml:"name"`
	Desc       string `yaml:"desc"`
	KubeConfig string `yaml:"kubeConfig"`
}

var Usage = func() {
	flagOutput := flag.CommandLine.Output()
	fmt.Fprintln(flagOutput, "Usage: realm [OPTIONS] COMMAND")
	fmt.Fprintln(flagOutput, "realms management of kbd, one realm is one k8s context")
	fmt.Fprintln(flagOutput, "Options:")
	fmt.Fprintln(flagOutput, "Commands:")
	fmt.Fprintln(flagOutput, "	set	create or update a realm")
	fmt.Fprintln(flagOutput, "	rm	remove a realm")
	fmt.Fprintln(flagOutput, "	use	active the realm")
}

func CreateRealmFlagSet(args []string) {
	fs := flag.NewFlagSet("realm", flag.ExitOnError)

	if len(args) == 0 {
		Usage()
		return
	}
	action := args[0]
	switch action {
	case "rm":
		fs.StringVar(&ram.Name, "name", "", "name of the realm")
		break
	case "set":
		fs.StringVar(&ram.KubeConfig, "name", "", "name of the realm")
		fs.StringVar(&ram.Desc, "desc", "", "desc of the realm")
		fs.StringVar(&ram.KubeConfig, "kubeconfig", "", "kubeconfig path of the k8s cluster")
		break
	case "use":
		fs.StringVar(&ram.KubeConfig, "name", "", "name of the realm")
		break
	default:
		Usage()
	}
	fs.Parse(args[1:])
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
