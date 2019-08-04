package sshClient

import (
	"sftp/pwdfilgen"
	"testing"

	"golang.org/x/crypto/ssh"
)

var inp = []struct {
	inp      string
	pwd      string
	protocol string
	host     string
	exp      ssh.Session
}{
	{"test", "tester", "tcp", "localhost:22", ssh.Session},
	{"test1", "tester1", "tcp", "localhost:22", ssh.Session},
	{"test2", "tester2", "tcp", "localhost:22", ssh.Session},
}

func TestSSHSession(t *testing.T) {
	//setstring up uname,pwd as inp val
	for _, v := range inp {
		uname := v.inp
		pwd := pwdfilgen.DecodePwd(v.pwd)
		exp := v.exp
	}

}
