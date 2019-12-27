package main

import (
	"log"
	"os"
	"sftp/helpcontent"
	inp "sftp/input"
	pwdfilgen "sftp/pwdfilgen"
	"sftp/sshClient"
)

/* main operations */
func main() {
	inpval := os.Args[1:]
	if len(inpval) < 1 {
		log.Print(helpcontent.Helpcontent())
	} else {
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
			inp.SetUserInp(string(usr), string(pwd), string(url), string(srcdir), string(tgtdir))
			log.Printf("TRACE:\t starting conn for %s", inp.Inp.Uname)
			log.Printf("TRACE:\t Password acquired for user  %s from %s", string(usr), pwdfilgen.Filnam)
			log.Printf("TRACE:\t attempting   connection")
			sshClient.SSHSession(string(usr), string(pwd))
			log.Printf("TRACE:\t starting conn for %s", inp.Inp.Url)
			log.Printf("TRACE:\t starting conn for %s", inp.Inp.SrcDir)
			log.Printf("TRACE:\t starting conn for %s", inp.Inp.TgtDir)
			// sshFilOperation.sshFilOpr() file operation in the new func

		default:
			log.Print(helpcontent.Helpcontent())
		}

	}

}
