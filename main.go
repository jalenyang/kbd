package main

import (
	"flag"
	"fmt"
)

var (
	kbdConfig KbdConfig
)

var (
	realm     string
	save      bool
	install   bool
	uninstall bool
	upgrade   bool
	rollback  bool
	use       string
)

var Usage = func() {
	fmt.Fprintln(flag.CommandLine.Output(), "Usage:  kbd [OPTIONS] COMMAND")
	fmt.Fprintln(flag.CommandLine.Output(), "kubernetes based development toolkit")
	fmt.Fprintln(flag.CommandLine.Output(), "Options:")
	fmt.Fprintln(flag.CommandLine.Output(), "	-D	enable debug mode")
	fmt.Fprintln(flag.CommandLine.Output(), "Commands:")
	fmt.Fprintln(flag.CommandLine.Output(), "	realm 	realm management of the kbd")
	fmt.Fprintln(flag.CommandLine.Output(), "	use 	active the realm")
	fmt.Fprintln(flag.CommandLine.Output(), "	helm 	toolset support of the helm client")
}

func init() {
	flag.StringVar(&realm, "realm", "", "The realm of the k8s cluster")
	flag.StringVar(&use, "use", "", "set realm to be used")
	flag.Usage = Usage
}

func main() {
	flag.Parse()
	action := flag.Arg(0)

	switch action {
	case "helm":
		helmClient := HelmClient{}
		loadFromConfigFile("kbd.yaml")
		helmClient.execHelmCommand(flag.Args()[1:]...)
	default:
		Usage()
	}

}
