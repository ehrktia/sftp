package main

import (
	"log"
	"os"

	"github.com/ehrktia/sftp/cmd"
)

func main() {
	log := log.New(os.Stderr, "Msg:", log.Ldate|log.Ltime|log.Lshortfile)
	inpval := os.Args[1:]
	cmd.CheckOption(log, inpval)
	// if !cmd.IsValidArgs(inpval) {
		// os.Exit(0)
	// }
}
