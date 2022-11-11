package sshClient

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/ssh"
)

var ErrEmptyConfig error = fmt.Errorf("%s", "empty config")

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

func validateHost(host string) bool {
	return strings.Contains(host, ":")

}

func NewSSHConfig(username, password, host string) sshConfig {
	if !validateInput(username, password, host) {
		return sshConfig{}
	}
	if !validateHost(host) {
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

func validateConfig(sshConfig sshConfig) bool {
	return strings.EqualFold(sshConfig.host, "")
}

func NewSSHClient(sshConfig sshConfig) (*ssh.Client, error) {
	if validateConfig(sshConfig) {
		return nil, ErrEmptyConfig
	}
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
