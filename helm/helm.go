package helm

import (
	"context"
	"kbd/realm"
	"log"
	"os/exec"
	"time"
)

type Client struct {
}

func (h Client) Install() error {
	return nil
}

func (h Client) Uninstall() error {
	return nil
}

func (h Client) Upgrade() error {
	return nil
}

func (h Client) Rollback() error {
	return nil
}

func (h Client) ExecHelmCommand(args []string) error {
	realm, err := realm.GetTheActiveRealm()
	if err != nil {
		log.Printf("Failed to get the active realm %v", err)
	}
	args = append(args, "--kubeconfig", realm.KubeConfig)
	_, err = exec.LookPath("helm")
	if err != nil {
		log.Printf("Warning: helm is not available on your host")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
	defer cancel()

	cmd := exec.CommandContext(ctx, "helm", args...)
	if output, err := cmd.Output(); err != nil {
		log.Printf("Run cmd helm failed for %v", err)
		return err
	} else {
		log.Println(string(output))
	}
	return nil
}
