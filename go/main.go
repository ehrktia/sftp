package main

import (
	"log"
	"os"

	"github.com/ehrktia/sftp/cmd"
)

func main() {
	log := log.New(os.Stderr, "Msg:", log.Ldate|log.Ltime|log.Lshortfile)
	inpval := os.Args[1:]
	if len(inpval) < 1 {
		log.Fatalf("%s\n", cmd.HelpMsg)
	}
	cmd.CheckOption(log, inpval)
}
