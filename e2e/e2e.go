package e2e

import (
	"flag"
	"fmt"
)

var e2e E2e

type E2e struct {
	pabotArgs string
	robotArgs string
	processes int
	rpEnabled bool
}

var Usage = func() {
	flagOutput := flag.CommandLine.Output()
	fmt.Fprintln(flagOutput, "Usage: e2e [OPTIONS] COMMAND")
	fmt.Fprintln(flagOutput, "kubernetes based e2e verification toolkit")
	fmt.Fprintln(flagOutput, "Options:")
	fmt.Fprintln(flagOutput, "Commands:")
	fmt.Fprintln(flagOutput, "	start 	start the e2e test on the k8s cluster")
	fmt.Fprintln(flagOutput, "	status 	check the status of the e2e test")
	fmt.Fprintln(flagOutput, "	kill 	kill the e2e test on the k8s cluster")
	fmt.Fprintln(flagOutput, "	results get the results of the e2e test")
}

func CreateE2eFlagSet(args []string) {
	fs := flag.NewFlagSet("e2e", flag.ExitOnError)
	if len(args) == 0 {
		Usage()
		return
	}
	action := args[0]
	switch action {
	case "start":
		fs.StringVar(&e2e.pabotArgs, "pabotargs", "", "Additional arguments for pabot command.")
		fs.StringVar(&e2e.robotArgs, "robotargs", "", "Additional arguments for robot command.")
		fs.IntVar(&e2e.processes, "processes", 0, "How many parallel executors to use with pabot.")
		fs.BoolVar(&e2e.rpEnabled, "rp-enabled", true, "Whether to use Report Portal or not")
		break
	case "status":
		break
	case "kill":
		break
	case "results":
		break
	default:
		Usage()
	}
	fs.Parse(args[1:])
}
