package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "porigo",
	Short: "fast open port checker",
	Long:  "porigo is a lightweight CLI tool to quickly check whether a port is open",
	Run: func(cmd *cobra.Command, args []string) {
		if Version {
			fmt.Println("porigo v1.0")
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var Version bool

func init() {
	rootCmd.Flags().BoolVarP(&Version, "version", "v", false, "show version")
}
