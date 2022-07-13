package main

import (
	"flag"
	"fmt"
	"kbd/e2e"
	"kbd/helm"
	rlm "kbd/realm"
	"log"
)

var Usage = func() {
	flagOutput := flag.CommandLine.Output()
	fmt.Fprintln(flagOutput, "Usage:  kbd [OPTIONS] COMMAND")
	fmt.Fprintln(flagOutput, "kubernetes based development toolkit")
	fmt.Fprintln(flagOutput, "Options:")
	fmt.Fprintln(flagOutput, "	-d	enable debug mode")
	fmt.Fprintln(flagOutput, "Commands:")
	fmt.Fprintln(flagOutput, "	e2e 	manage the e2e tests in the kubernetes cluster")
	fmt.Fprintln(flagOutput, "	realm 	manage the realms in the kbd")
	fmt.Fprintln(flagOutput, "	helm 	manage the helm business in the kubernetes cluster")
}

func init() {
	flag.Usage = Usage
}

func main() {
	flag.Parse()
	action := flag.Arg(0)

	switch action {
	case "e2e":
		e2e.CreateE2eFlagSet(flag.Args()[1:])
		break
	case "helm":
		helmClient := helm.HelmClient{}
		helmClient.ExecHelmCommand(flag.Args()[1:]...)
		break
	case "realm":
		if err := rlm.Operator(flag.Args()[1:]); err != nil {
			log.Printf("realm failed for %v", err)
		}
		break
	default:
		flag.Usage()
	}

}
