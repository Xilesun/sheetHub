package commands

import (
	"context"

	"github.com/spf13/cobra"
)

func init() {
	appCmd.AddCommand(installCmd)
}

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install sheethub application",
	Run: func(cmd *cobra.Command, args []string) {
		sheethub.Install(context.Background())
	},
}
