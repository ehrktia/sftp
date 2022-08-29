package cmd

import (
	"fmt"
	"log"

	"github.com/ehrktia/sftp/helpcontent"
	"github.com/ehrktia/sftp/input"
	"github.com/ehrktia/sftp/pwdfilgen"
	"github.com/ehrktia/sftp/sshClient"
)

const (
	help       = "help"
	genPwdFile = "GenPwdFile"
	conn       = "conn"
	invalid    = "default"
)

// IsValidArgs checks length of args from input stdin
func IsValidArgs(args []string) bool {
	return len(args) > 1
}

func CheckOption(l *log.Logger, args []string) {
	switch string(args[0]) {
	case help:
		l.Print(helpcontent.Helpcontent(l))
	case genPwdFile:
		genpwd := []byte(args[1])
		l.Printf("%s-%v\n", "INFO", genpwd)
		if len(string(genpwd)) < 7 {
			l.Fatalf("%s-%s\n", "ERROR", "min length for pwd is 7 characters")
		}
		if err := pwdfilgen.GenPwdFile(genpwd); err != nil {
			panic(fmt.Sprintf("error generating pwd:%v\n", err))
		}
	case conn:
		usr := []byte(args[1])
		pwd := pwdfilgen.DecodePwd()
		url := []byte(args[2])
		srcdir := []byte(args[3])
		tgtdir := []byte(args[4])
		input.SetUserInp(string(usr), string(pwd), string(url), string(srcdir), string(tgtdir))
		l.Printf("%s-%s\n", "INFO", input.Inp.Uname)
		l.Printf("%s-Password acquired for user:%s from %s\n", "INFO", string(usr), pwdfilgen.Filnam)
		l.Printf("%s-%s\n", "INFO", "attempting connection")
		sshClient.SSHSession(string(usr), string(pwd))
		l.Printf("%s-starting conn for %s", "INFO", input.Inp.Url)
		l.Printf("%s-starting conn for %s", "INFO", input.Inp.SrcDir)
		l.Printf("%s-starting conn for %s", "INFO", input.Inp.TgtDir)
	case invalid:
		l.Print(helpcontent.Helpcontent(l))
	}
}
