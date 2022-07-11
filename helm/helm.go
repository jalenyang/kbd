package helm

import (
	"context"
	"log"
	"os/exec"
	"time"
)

type HelmClient struct {
}

func (h HelmClient) Install() error {
	return nil
}

func (h HelmClient) Uninstall() error {
	return nil
}

func (h HelmClient) Upgrade() error {
	return nil
}

func (h HelmClient) Rollback() error {
	return nil
}

func (h HelmClient) ExecHelmCommand(args ...string) error {
	_, err := exec.LookPath("helm")
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
