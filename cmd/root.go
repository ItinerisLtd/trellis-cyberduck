package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// To be replaced by goreleaser build flag.
var version = "canary"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "trellis-cyberduck",
	Version: version,
	Short: "Trellis commands for Cyberduck",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
