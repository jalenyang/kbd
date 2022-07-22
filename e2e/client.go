package e2e

import (
	"errors"
	"fmt"
	"io"
	"kbd/realm"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type Client struct {
	e2eApi string
}

func CreateFromKbdConfig() (Client, error) {
	realm, err := realm.GetTheActiveRealm()
	if err != nil {
		log.Printf("Failed create the e2e client: %v", err)
		return Client{}, err
	}
	return Client{e2eApi: realm.E2eTestApi}, nil
}

func (c *Client) renderApi(apiName string) string {
	return fmt.Sprintf("%s/%s", c.e2eApi, apiName)
}

func (c *Client) Start(values E2e) error {
	data := fmt.Sprintf(`--pabotargs="%s" --robotargs="%s" --processes %d  --rp_enabled=%t`, values.pabotArgs,
		values.robotArgs, values.processes, values.rpEnabled)
	payload := strings.NewReader(data)

	client := http.Client{}
	req, err := http.NewRequest(http.MethodPost, c.renderApi("start"), payload)
	if err != nil {
		fmt.Fprintln(os.Stdout, fmt.Sprintf("create the request failed for %v", err))
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stdout, fmt.Sprintf("send the request failed for %v", err))
		return err
	}
	if res.StatusCode != 200 {
		fmt.Fprintln(os.Stdout, fmt.Sprintf("start the e2e test failed with status: %s", res.Status))
		return errors.New(fmt.Sprintf("start the e2e test failed with status: %s", res.Status))
	}
	fmt.Fprintln(os.Stdout, "start the e2e test successfully")
	return nil
}

func (c *Client) Status() error {
	res, err := http.Get(c.renderApi("status"))
	if err != nil {
		fmt.Fprintln(os.Stdout, fmt.Sprintf("send the request failed for %v", err))
		return err
	}
	if res.StatusCode != 200 {
		fmt.Fprintln(os.Stdout, fmt.Sprintf("query the status failed with: %s", res.Status))
		return errors.New(fmt.Sprintf("query the status failed with: %s", res.Status))
	}
	defer res.Body.Close()
	result, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Fprintln(os.Stdout, fmt.Sprintf("Failed to read the body for: %v", err))
		return err
	}
	fmt.Fprintln(os.Stdout, fmt.Sprintf("the e2e execution status is %s", string(result)))
	return err
}

func (c *Client) Kill() error {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodDelete, c.renderApi("kill"), nil)
	if err != nil {
		fmt.Fprintln(os.Stdout, fmt.Sprintf("create the request failed for %v", err))
		return err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stdout, fmt.Sprintf("send the request failed for %v", err))
		return err
	}
	if res.StatusCode != 200 {
		fmt.Fprintln(os.Stdout, fmt.Sprintf("failed to kill the e2e cases for: %v", err))
		return err
	}
	fmt.Fprintln(os.Stdout, "kill the e2e cases successful")
	return nil
}

func (c *Client) Results() error {
	res, err := http.Get(c.renderApi("results.zip"))
	if err != nil {
		fmt.Fprintln(os.Stdout, fmt.Sprintf("send the request failed for %v", err))
		return err
	}
	defer res.Body.Close()
	results, err := io.ReadAll(res.Body)
	targetDir := path.Join(os.TempDir(), "results.zip")

	fi, err := os.OpenFile(targetDir, os.O_RDWR|os.O_CREATE, 0655)
	if err != nil {
		log.Panicln("Error: failed to save the results file")
	}
	_, err = fi.Write(results)
	if err != nil {
		log.Panicln("Error: failed to save the results file")
	}
	fmt.Fprintln(os.Stdout, fmt.Sprintf("saved the results package to %s", targetDir))
	return err
}
