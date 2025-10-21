package main

import (
	"fmt"

	"github.com/bitterpanda63/michi/tui"
	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
)

var rootCmd = &cobra.Command{
	Use:   "michi",
	Short: "A TUI for interacting with AI models",
	Run: func(cmd *cobra.Command, args []string) {
		tui.RunTUI()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of michi",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("michi v%s\n", version)
	},
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate with your API token",
	Run: func(cmd *cobra.Command, args []string) {
		tui.RunAuthTUI()
	},
}

func main() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(authCmd)
	rootCmd.Execute()
}
