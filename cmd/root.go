package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/xiaobaiskill/blockchain-release-monitor/cmd/server"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "block-chain-release",
	Short: "",
	Long:  ``,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(
		server.NewServer(),
	)

	rootCmd.PersistentFlags().BoolP("help", "h", false, "help")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
