package cmd

import (
	"porigo/function"

	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "quickly check if a port is open",
	Long:  "a quick scan checks whether ports are open on the target host.",
	Run: func(cmd *cobra.Command, args []string) {
		myport := function.ParseFlag(Port)
		function.ScannerPort(myport, Target, Download)
	},
}

var Port string
var Target string
var Download bool

func init() {
	rootCmd.AddCommand(scanCmd)
	scanCmd.Flags().StringVarP(&Port, "port", "p", "1-1000", "use format 1-1024 or 80")
	scanCmd.Flags().StringVarP(&Target, "target", "t", "google.com", "set target")
	scanCmd.Flags().BoolVarP(&Download, "download", "d", false, "download result")

}
