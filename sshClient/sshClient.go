package sshClient

import (
	"log"

	"golang.org/x/crypto/ssh"
)

type sshConfig struct {
	sshClientconfig *ssh.ClientConfig
	host            string
}

func validateInput(username, password, host string) bool {
	if username == "" || password == "" || host == "" {
		return false
	}
	return true
}

func NewSSHConfig(username, password, host string) sshConfig {
	if !validateInput(username, password, host) {
		return sshConfig{}
	}
	return sshConfig{
		sshClientconfig: &ssh.ClientConfig{
			User: username,
			Auth: []ssh.AuthMethod{
				ssh.Password(password),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		},
		host: host,
	}

}

func NewSSHClient(sshConfig sshConfig) (*ssh.Client, error) {
	client, err := ssh.Dial("tcp", sshConfig.host, sshConfig.sshClientconfig)
	if err != nil {
		return nil, err
	}
	return client, nil

}

func SSHSession(l *log.Logger, client *ssh.Client) (*ssh.Session, error) {

	// var hostKey ssh.PublicKey
	// An SSH client is represented with a ClientConn.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig,
	// and provide a HostKeyCallback.

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
