package k8s

import (
	"context"
	"fmt"
	"kbd/realm"
	"log"
	"os"
	"os/exec"
	"time"
)

func ExecK8sCommand(args []string) error {
	realm, err := realm.GetTheActiveRealm()
	if err != nil {
		log.Printf("Failed to get the active realm %v", err)
	}
	defaultArgs := []string{"--kubeconfig", realm.KubeConfig}
	args = append(defaultArgs, args...)

	kClient := "oc"
	// first try oc
	_, err = exec.LookPath("oc")
	if err != nil {
		fmt.Fprintln(os.Stdout, "Warning: oc is not available on your host, will try kubectl")
		// secondly try kubectl if oc not found
		_, err = exec.LookPath("oc")
		if err != nil {
			fmt.Fprintln(os.Stdout, "Error: kubectl is not available on your host, exit now")
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()
	cmd := exec.CommandContext(ctx, kClient, args...)
	if output, err := cmd.Output(); err != nil {
		fmt.Fprintln(os.Stdout, fmt.Sprintf("Run cmd %s failed for %v", kClient, err))
		return err
	} else {
		fmt.Fprintln(os.Stdout, string(output))
	}
	return nil
}
