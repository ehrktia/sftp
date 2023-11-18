package cmd

import (
	"os"
	"testing"
)

func Test_cmd(t *testing.T) {
	t.Run("run help command successfully", func(t *testing.T) {
		os.Args = []string{"help"}
		testHelpCommand := &HelpCommand{Name: t.Name()}
		want := testHelpCommand.Run(os.Args)
		t.Log(want)
		t.Log("completed")
	})

}
