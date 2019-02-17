package main

import (
	"flag"
	"log"
	"os"

	inp "github.com/sftp/input"
	pwdfilgen "github.com/sftp/pwdfilgen"
	"github.com/sftp/sshClient"

	"github.com/sftp/helpcontent"
)

var con = flag.NewFlagSet("conn", flag.ExitOnError)
var pwd = flag.NewFlagSet("pwdgen", flag.ExitOnError)
var csr = con.String("serv", "", "name of server to connect with")
var cun = con.String("uname", "", "uname to server to connect with")
var srd = con.String("src", "", "src dir to move files from")
var tgt = con.String("tgt", "", "tgt dir to mov file to")
var cmd = con.String("cmd", "", "command which is required to be executed")
var pval = pwd.String("pwd", "", "provide pwd for the user")

/* main operations */
func main() {
	if len(os.Args) < 2 {
		log.Print(helpcontent.Helpcontent())
		log.Fatalf("ERROR:\tno valid arg provided")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "help":
		log.Print(helpcontent.Helpcontent())
	case "pwdgen":
		pwd.Parse(os.Args[2:])
	case "conn":
		con.Parse(os.Args[2:])
	default:
		log.Print(helpcontent.Helpcontent())
	}
	if con.Parsed() {
		if *csr == "" || *cun == "" || *srd == "" || *tgt == "" {
			log.Fatal("ERROR:\tcannot take empty arg for conn parameter")
			os.Exit(1)
		} else {
			usr := *cun
			pwd := pwdfilgen.DecodePwd()
			url := *csr
			cmd := *cmd
			srcdir := *srd
			tgtdir := *tgt
			inp.SetUserInp(string(usr), string(pwd), string(url), string(srcdir), string(tgtdir))
			log.Printf("TRACE:\tstarting conn for %s", inp.Inp.Uname)
			log.Printf("TRACE:\tPassword acquired for user  %s from %s", string(usr), pwdfilgen.Filnam)
			log.Printf("TRACE:\tattempting   connection")
			sshClient.SshSession(string(usr), string(pwd), url, cmd, srcdir, tgtdir)
			log.Printf("TRACE:\tstarting conn for server  %s", inp.Inp.Url)
			log.Printf("TRACE:\tcopying from  %s", inp.Inp.SrcDir)
			log.Printf("TRACE:\tdestination dir used is %s", inp.Inp.TgtDir)
		}

	}
	if pwd.Parsed() {
		if *pval == "" {
			log.Fatal("ERROR:\tcannot take empty arg for pwd parameter")
			os.Exit(1)
		} else {

			genpwd := []byte(*pval)
			if len(string(genpwd)) < 7 {
				log.Fatal("ERROR:\tmin length for pwd is 7 characters")
			}
			pwdfilgen.GenPwdFile(genpwd)
		}
	}
}
