package vscode

import (
	"os/exec"
)

func checkDocker() error {
	cmd := exec.Command("docker -v")
	return cmd.Start()
}

// attach to docker container
func upContainer() {}

func getContaierInfo(name string) string {
	containerName := "{\"containerName\":\"" + name + "\"}"
	return convertHex(containerName)
}
