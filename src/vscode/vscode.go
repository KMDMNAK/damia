package vscode

import (
	"errors"
	"os/exec"

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
	cmd := genCommand(args[0], flags)
	err = cmd.Start()
	return err
}

func genCommand(arg string, flags *models.VscodeOpenFlags) *exec.Cmd {
	var cmd *exec.Cmd
	if flags.Type == "dir" {
		dirpath := arg
		cmd = exec.Command("code", dirpath)
	} else if flags.Type == "service" {
		// TODO
		cmd = exec.Command("code")
	} else if flags.Type == "container" {
		containername := arg
		cmd = exec.Command("code", "--folder-uri=vscode-remote://attached-container+"+getContaierInfo(containername)+flags.WORK_DIR)
	}
	return cmd
}
