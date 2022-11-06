package sshClient

import (
	"log"

	"golang.org/x/crypto/ssh"
)

func SSHSession(l *log.Logger, uname, pwd, host string) (*ssh.Session, error) {

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
	client, err := ssh.Dial("tcp", host, config)
	if err != nil {
		return nil, err
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	defer session.Close()
	l.Printf("INFO:  session started")
	return session, nil

	/* 	// Once a Session is created, you can execute a single command on
	   	// the remote side using the Run method.
	   	var b bytes.Buffer
	   	session.Stdout = &b
	   	if err := session.Run("date"); err != nil {
	   		log.Fatal("Failed to run: " + err.Error())
	   	}
	   	fmt.Println(b.String())
	*/
}
