package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ehrktia/sftp/cmd"
	"github.com/ehrktia/sftp/helpcontent"
	"github.com/ehrktia/sftp/input"
	"github.com/ehrktia/sftp/pwdfilgen"
	"github.com/ehrktia/sftp/sshClient"
)

func main() {
	inpval := os.Args[1:]
	if !cmd.IsValidArgs(inpval) {
		panic(fmt.Sprintf("%s-%s", "ERR", "missing required number of args"))

	}
	switch string(inpval[0]) {
	case "help":
		log.Print(helpcontent.Helpcontent())
	case "GenPwdFile":
		genpwd := []byte(inpval[1])
		log.Print(genpwd)
		if len(string(genpwd)) < 7 {
			log.Fatal("ERROR: min length for pwd is 7 characters")

		}

		pwdfilgen.GenPwdFile(genpwd)
	case "conn":
		usr := []byte(inpval[1])
		pwd := pwdfilgen.DecodePwd()
		url := []byte(inpval[2])
		srcdir := []byte(inpval[3])
		tgtdir := []byte(inpval[4])
		input.SetUserInp(string(usr), string(pwd), string(url), string(srcdir), string(tgtdir))
		log.Printf("TRACE:\t starting conn for %s", input.Inp.Uname)
		log.Printf("TRACE:\t Password acquired for user  %s from %s", string(usr), pwdfilgen.Filnam)
		log.Printf("TRACE:\t attempting   connection")
		sshClient.SSHSession(string(usr), string(pwd))
		log.Printf("TRACE:\t starting conn for %s", input.Inp.Url)
		log.Printf("TRACE:\t starting conn for %s", input.Inp.SrcDir)
		log.Printf("TRACE:\t starting conn for %s", input.Inp.TgtDir)

	default:
		log.Print(helpcontent.Helpcontent())
	}

}
