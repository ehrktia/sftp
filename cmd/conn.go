package cmd

import (
	"flag"
	"io"
	"os"
)

type connCommand struct {
	output io.Writer
	name   string
}

func (h *connCommand) Help() string {
	return "am conn"

}

func (h *connCommand) Run(args []string) int {
	if h.output == nil {
		h.output = os.Stdout
	}
	f := flag.NewFlagSet(conn, flag.ContinueOnError)
	f.StringVar(&h.name, conn, "conn", "conn message to be displayed")
	_, _ = h.output.Write([]byte(h.Help()))
	if err := f.Parse(args); err != nil {
		return 1
	}
	return 0
}

func (h *connCommand) Synopsis() string {
	return "conn command list all available option"

}
