package cmd

import (
	"flag"
	"io"
	"os"
)

type helpCommand struct {
	output io.Writer
	name   string
}

func (h *helpCommand) Help() string {
	return "am help"

}

func (h *helpCommand) Run(args []string) int {
	if h.output == nil {
		h.output = os.Stdout
	}
	f := flag.NewFlagSet(help, flag.ContinueOnError)
	f.StringVar(&h.name, help, "help", "help message to be displayed")
	_, _ = h.output.Write([]byte(h.Help()))
	if err := f.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (h *helpCommand) Synopsis() string {
	return "help command list all available option"

}
