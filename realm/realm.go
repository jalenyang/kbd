package realm

import (
	"errors"
	"flag"
	"fmt"
	"kbd/config"
	"kbd/module"
	"kbd/util"
	"log"
)

var (
	ram       module.Realm
	kbdConfig module.KbdConfig
)

var Usage = func() {
	flagOutput := flag.CommandLine.Output()
	fmt.Fprintln(flagOutput, "Usage: realm [OPTIONS] COMMAND")
	fmt.Fprintln(flagOutput, "realms management of kbd, one realm is one k8s context")
	fmt.Fprintln(flagOutput, "Options:")
	fmt.Fprintln(flagOutput, "Commands:")
	fmt.Fprintln(flagOutput, "	ls	list the realms")
	fmt.Fprintln(flagOutput, "	rm	remove a realm")
	fmt.Fprintln(flagOutput, "	set	create or update a realm")
	fmt.Fprintln(flagOutput, "	use	active the realm")
}

func init() {
	var err error
	kbdConfig, err = config.LoadFromConfigFile()
	if err != nil {
		panic("Failed to parse the configuration file")
	}
}

func Operator(args []string) error {
	fs := flag.NewFlagSet("realm", flag.ExitOnError)
	if len(args) == 0 {
		Usage()
		return nil
	}
	action := args[0]
	switch action {
	case "ls":
		fs.String("ls", "all", "list the realms")
		listRealms()
	case "rm":
		fs.StringVar(&ram.Name, "name", "", "name of the realm")
		fs.Parse(args[1:])
		return rmRealm(ram)
	case "set":
		fs.StringVar(&ram.Name, "name", "", "name of the realm")
		fs.StringVar(&ram.Desc, "desc", "", "desc of the realm")
		fs.StringVar(&ram.KubeConfig, "kubeconfig", "", "kubeconfig path of the k8s cluster")
		fs.Parse(args[1:])
		return setRealm(ram)
	case "use":
		fs.StringVar(&ram.Name, "name", "", "name of the realm")
		fs.Parse(args[1:])
		return useRealm(ram)
	default:
		Usage()
	}
	return nil
}

func GetTheActiveRealm() (module.Realm, error) {
	kbdConfig, err := config.LoadFromConfigFile()
	if err != nil {
		log.Printf("Error: configuration error %v", kbdConfig)
		return module.Realm{}, err
	}
	for _, realm := range kbdConfig.Realms {
		if realm.Active {
			return realm, nil
		}
	}
	return module.Realm{}, nil
}

func listRealms() error {
	for _, realm := range kbdConfig.Realms {
		fmt.Println("Available realms:")
		if realm.Active {
			fmt.Println("*" + realm.Name)
		} else {
			fmt.Println(realm.Name)
		}
	}
	return nil
}

func setRealm(realm module.Realm) error {
	if realm.Name == "" || realm.KubeConfig == "" {
		log.Println("invalid realm, realm name and kubeconfig can't be empty")
		return errors.New("invalid realm, realm name and kubeconfig can't be empty")
	}
	if util.Contains(kbdConfig.Realms, realm) {
		index := util.Index(kbdConfig.Realms, realm)
		kbdConfig.Realms[index] = realm
	} else {
		realms := append(kbdConfig.Realms, realm)
		kbdConfig.Realms = realms
	}
	return config.SyncToConfigFile(kbdConfig)
}

func useRealm(realm module.Realm) error {
	if realm.Name == "" {
		log.Println("please specify the realm name")
	}
	if util.Contains(kbdConfig.Realms, realm) {
		index := util.Index(kbdConfig.Realms, realm)
		kbdConfig.Realms[index].Active = true
		return config.SyncToConfigFile(kbdConfig)
	}
	return errors.New("the realm doesn't exists")
}

func rmRealm(realm module.Realm) error {
	if realm.Name == "" {
		log.Println("please specify the realm name")
	}
	if util.Contains(kbdConfig.Realms, realm) {
		index := util.Index(kbdConfig.Realms, realm)
		kbdConfig.Realms = append(kbdConfig.Realms[:index], kbdConfig.Realms[index+1:]...)
		return config.SyncToConfigFile(kbdConfig)
	}
	return errors.New("the realm doesn't exists")
}
