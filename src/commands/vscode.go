package commands

import (
	"github.com/damia/models"
	"github.com/damia/vscode"
	"github.com/spf13/cobra"
)

func genVscodeCommand() *baseCmd {
	ccmd := &cobra.Command{
		Use: "vscode",
	}
	cc := newBaseCmd(ccmd)
	openCommand := genVscodeOpenCommand()
	cc.cmd.AddCommand(openCommand)
	return cc
}

func genVscodeOpenCommand() *cobra.Command {

	flags := &models.VscodeOpenFlags{
		Type:     "",
		Path:     "",
		WORK_DIR: "",
	}

	cmd := &cobra.Command{
		Use: "open",
		Run: func(cmd *cobra.Command, args []string) {
			err := vscode.RunOpen(args, flags)
			if err != nil {
				panic(err)
			}
		},
	}
	cmd.Flags().StringVarP(&flags.Type, "type", "t", "dir", "")
	cmd.Flags().StringVarP(&flags.Path, "path", "p", "", "")
	cmd.Flags().StringVarP(&flags.WORK_DIR, "work_dir", "w", "/root/workspace", "")

	return cmd
}
