package cmd

import (
	"flag"
	"io"
	"os"
)

var HelpMsg = `this is a command line operated program 
		To get list of available option  provide the command option -help=help 
		ex: uname=TE123455 pwd=password
		To generate key use key mode`

type HelpCommand struct {
	output io.Writer
	Name   string
}

func (h *HelpCommand) Help() string {
	return "am help"

}

func (h *HelpCommand) Run(args []string) int {
	if h.output == nil {
		h.output = os.Stdout
	}
	f := flag.NewFlagSet(help, flag.ContinueOnError)
	f.StringVar(&h.Name, help, "help", "help message to be displayed")
	_, _ = h.output.Write([]byte(h.Help()))
	if err := f.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (h *HelpCommand) Synopsis() string {
	return "help command list all available option"

}
