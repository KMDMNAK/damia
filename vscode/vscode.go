package vscode

import (
	"fmt"

	"github.com/damia/models"
)

func RunOpen(flags *models.VscodeOpenFlags) error {
	fmt.Println("This is Vscode Open")
	return nil
}
