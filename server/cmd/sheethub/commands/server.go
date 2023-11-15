package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	appCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntP("port", "p", 4000, "Port to listen on")
	viper.BindPFlag("app.port", serverCmd.Flags().Lookup("port"))
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Long:  "Start the sheethub server",
	Run: func(cmd *cobra.Command, args []string) {
		sheethub.Start()
	},
}
