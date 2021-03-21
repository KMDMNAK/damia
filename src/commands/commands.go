package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

type baseCmd struct {
	cmd *cobra.Command
}

func (c *baseCmd) getCommand() *cobra.Command {
	return c.cmd
}

func newBaseCmd(cmd *cobra.Command) *baseCmd {
	return &baseCmd{cmd: cmd}
}

func initCommands() *baseCmd {
	cmd := &cobra.Command{
		Use:   "damia",
		Short: "Damia is a command line tool for assistant development ",
		Long:  `you can collaborate with vscode, docker`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello Damia!")
		},
	}
	cmd.AddCommand(genVscodeCommand().getCommand())
	var rootCmd = newBaseCmd(cmd)
	return rootCmd
}

func Execute() error {
	cc := initCommands()
	return cc.getCommand().Execute()
}
