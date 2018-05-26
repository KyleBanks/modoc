package cmd

import (
	"fmt"

	"github.com/KyleBanks/modoc/pkg/document"
	"github.com/KyleBanks/modoc/pkg/document/markdown"
	"github.com/spf13/cobra"
)

type initCommand struct {
	*cobra.Command

	path string
}

func (i *initCommand) validate(cmd *cobra.Command, args []string) {
	path, err := cleanPath(i.path)
	if err != nil {
		fail("", err)
	}

	i.path = path
}

func (i *initCommand) run(cmd *cobra.Command, args []string) {
	fmt.Println("Generating source tree at", i.path)

	o := markdown.NewOrganizer(cmdOutput)
	if err := document.Sample.Organize(o, i.path); err != nil {
		fail("Failed to initialize source:", err)
	}
}

func init() {
	var i initCommand
	i.Command = &cobra.Command{
		Use:    "init",
		Short:  "Initializes a modoc project by creating a sample source tree.",
		PreRun: i.validate,
		Run:    i.run,
	}

	i.Flags().StringVarP(&i.path, "path", "p", "./", "The path to the root directory where the source tree will be generated.")
	rootCmd.AddCommand(i.Command)
}
