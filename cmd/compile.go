package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/KyleBanks/modoc/pkg/document/markdown"
	"github.com/spf13/cobra"
)

const defaultOutputFilename = "COMPILED.md"

type compileCommand struct {
	*cobra.Command

	source string
	output string

	header bool
	toc    bool
	body   bool
}

func (c *compileCommand) validate(cmd *cobra.Command, args []string) {
	source, err := cleanPath(c.source)
	if err != nil {
		fail("", err)
	}
	output, err := cleanPath(c.output)
	if err != nil {
		fail("", err)
	}

	if fi, err := os.Stat(source); err != nil {
		fail("Source directory does not exist:", err)
	} else if !fi.IsDir() {
		fail("Source must point to a directory:", fmt.Errorf("%v is not a directory", source))
	}

	if fi, err := os.Stat(output); err != nil && !os.IsNotExist(err) {
		fail("Failed accessing output file:", err)
	} else if fi != nil && fi.IsDir() {
		output = filepath.Join(output, defaultOutputFilename)
	}

	c.source = source
	c.output = output
}

func (c *compileCommand) run(cmd *cobra.Command, args []string) {
	fmt.Println("Compiling", c.source, "to", c.output)

	d, err := markdown.NewDocument(c.source)
	if err != nil {
		fail("Failed loading source files:", err)
	}

	mc := markdown.Compiler{
		Header: c.header,
		TOC:    c.toc,
		Body:   c.body,
	}
	if err := d.Compile(mc, cmdOutput, c.output); err != nil {
		fail("Failed compiling source files:", err)
	}
}

func init() {
	var c compileCommand
	c.Command = &cobra.Command{
		Use:    "compile",
		Short:  "Compiles the source tree into a single file.",
		PreRun: c.validate,
		Run:    c.run,
	}

	c.Flags().StringVarP(&c.source, "source", "s", "./", "The path to the source tree")
	c.Flags().StringVarP(&c.output, "output", "o", "./", "The path to the compiled Markdown file")
	c.Flags().BoolVarP(&c.header, "header", "", true, "If true, generates the header of the book including the title and optional introduction")
	c.Flags().BoolVarP(&c.toc, "toc", "", true, "If true, generates a table of contents at the start of the generated document.")
	c.Flags().BoolVarP(&c.body, "body", "", true, "If true, the main content of the document is generated.")
	rootCmd.AddCommand(c.Command)
}
