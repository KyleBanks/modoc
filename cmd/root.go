package cmd

import (
	"fmt"
	"os"

	"github.com/KyleBanks/modoc/pkg/output"
	"github.com/spf13/cobra"
)

var cmdOutput output.Writer = output.FileSystem{}

var rootCmd = &cobra.Command{
	Use:   "modoc",
	Short: "modoc compiles a Markdown file from a source tree of partial components.",
}

// Execute runs the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
}
