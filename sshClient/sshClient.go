package sshClient

import (
	"bytes"
	"log"

	"golang.org/x/crypto/ssh"
)

func SshSession(uname, pwd, srv, cmd, sd, td string) {

	// var hostKey ssh.PublicKey
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.
	config := &ssh.ClientConfig{
		User: uname,
		Auth: []ssh.AuthMethod{
			ssh.Password(pwd),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	client, err := ssh.Dial("tcp", "localhost:22", config)
	if err != nil {
		log.Fatal("ERROR:\t Failed to dial: ", err)
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		log.Fatal("ERROR:\t Failed to create session: ", err)
	}
	defer session.Close()
	log.Printf("TRACE:\t session started")

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run(string(cmd) + " " + string(sd) + " " + string(td)); err != nil {
		log.Fatal("ERROR:\t Failed to run: " + err.Error())
	}

}
