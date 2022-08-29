package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ehrktia/sftp/cmd"
	"github.com/ehrktia/sftp/helpcontent"
)

func main() {
	log := log.New(os.Stderr, "Msg:", log.Ldate|log.Ltime|log.Lshortfile)
	inpval := os.Args[1:]
	if !cmd.IsValidArgs(inpval) {
		log.Print(helpcontent.Helpcontent(log))
		os.Exit(0)
	cmd.CheckOption(log, inpval)
}
