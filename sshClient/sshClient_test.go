package sshClient

import (
	"log"
	"testing"
)

func TestSSHSession(t *testing.T) {
	l := log.Default()
	_, err := SSHSession(l, t.Name(), t.Name(), "localhost:22")
	if err == nil {
		t.Fatal("expected to fail,got no error")
	}

}
