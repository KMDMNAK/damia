package vscode

import (
	"errors"
	"os/exec"

	"github.com/damia/docker"
	"github.com/damia/models"
)

const (
	DIRECTORY = iota
	COMPOSEFILE
	DOCKERFILE
	CONTAINERNAME
)

func valid_flags(args []string, flags *models.VscodeOpenFlags) error {
	if len(args) != 1 {
		return errors.New("invalid Args")
	}
	if flags.Type == "dir" {
		return nil
	} else if flags.Type == "service" {
		return nil
	} else if flags.Type == "container" {
		return nil
	}
	return errors.New("invalid Type Flag")
}

func RunOpen(args []string, flags *models.VscodeOpenFlags) error {
	err := valid_flags(args, flags)
	if err != nil {
		return err
	}
	cmd, err := genCommand(args[0], flags)
	if err != nil {
		return err
	}
	err = cmd.Start()
	return err
}

func genCommand(arg string, flags *models.VscodeOpenFlags) (*exec.Cmd, error) {
	var cmd *exec.Cmd
	var err error
	if flags.Type == "dir" {
		dirpath := arg
		cmd = exec.Command("code", dirpath)
	} else {
		var containername string
		if flags.Type == "service" {
			// TODO
			serviceName := arg
			composepath := flags.Path
			containername, err = docker.GetContainerName(serviceName, composepath)
			if err != nil {
				return nil, err
			}
		} else if flags.Type == "container" {
			containername = arg
		}
		cmd = exec.Command("code", "--folder-uri=vscode-remote://attached-container+"+getContaierInfo(containername)+flags.WORK_DIR)
	}
	return cmd, nil
}
