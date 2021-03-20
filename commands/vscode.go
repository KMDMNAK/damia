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
		C: "",
		D: "",
		S: "",
	}

	cmd := &cobra.Command{
		Use: "open",
		Run: func(cmd *cobra.Command, args []string) {
			err := vscode.RunOpen(flags)
			if err != nil {
				panic(err)
			}
		},
	}
	cmd.Flags().StringVarP(&flags.C, "compose-file", "c", "", "provide docker compose file path.\nYou have to provide service name setting --s option.")
	cmd.Flags().StringVarP(&flags.D, "docker-file", "d", "", "provide docker file path.\nOpen attached vscode windows.")
	cmd.Flags().StringVarP(&flags.S, "service-name", "s", "", "provide docker compose service name.")

	return cmd
}
