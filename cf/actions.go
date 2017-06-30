package cf

import (
	"fmt"
	"os/exec"
	"strings"
)

func Push(name, directory string) error {
	_, err := runCF("push", name, "-p", directory)
	return err
}

func Delete(appName string) error {
	_, err := runCF("delete", "-f", appName)
	return err
}

func AppGuid(name string) (string, error) {
	output, err := runCF("app", name, "--guid")
	return strings.TrimSpace(string(output)), err
}

func runCF(args ...string) (string, error) {
	fmt.Printf("running: cf %s\n", strings.Join(args, " "))
	output, err := exec.Command("cf", args...).CombinedOutput()
	if err != nil {
		fmt.Printf("error running above command. Output: '%s', error: '%s'\n", string(output), err)
		return "", err
	}
	return string(output), nil
}