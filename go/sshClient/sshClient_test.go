package sshClient

import (
	"testing"
)

func TestSSHSession(t *testing.T) {
	// l := log.Default()
	// _, err := SSHSession(l, t.Name(), t.Name(), "localhost:22")
	// if err == nil {
	// t.Fatal("expected to fail,got no error")
	// }
	//
}

func TestValidateInput(t *testing.T) {
	t.Run("return true when all inputs are valid", func(t *testing.T) {
		got := validateInput(t.Name(), t.Name(), t.Name())
		if !got {
			t.Fatal("expected to get `true`,got `false`")
		}
	})
	t.Run("return false when any of one inputs are empty", func(t *testing.T) {
		got := validateInput(t.Name(), "", t.Name())
		if got {
			t.Fatal("expected to get `false`,got `true`")
		}
	})
	t.Run("return false when any of one inputs are empty", func(t *testing.T) {
	})

}

func TestNewConfig(t *testing.T) {
	t.Run("generate new ssh config when usr pwd host provided", func(t *testing.T) {
		usr, pwd, host := t.Name(), t.Name(), t.Name()+":111"
		got := NewSSHConfig(usr, pwd, host)
		if got.sshClientconfig == nil {
			t.Fatal("error generating config")
		}
		validateGot := validateConfig(got)
		if validateGot {
			t.Fatal("expected to get valid config")
		}

	})
	t.Run("fail when empty input provided", func(t *testing.T) {
		got := NewSSHConfig(t.Name(), "", t.Name())
		if got.sshClientconfig != nil {
			t.Fatal("should not generate a valid config with empty input")
		}
	})
	t.Run("fail when host provided with missing port", func(t *testing.T) {
		got := NewSSHConfig(t.Name(), t.Name(), t.Name())
		if got.sshClientconfig != nil {
			t.Fatal("expected to get empty config")
		}
	})

}

func TestNewSSHClient(t *testing.T) {
	testSSHConfig := NewSSHConfig(t.Name(), t.Name(), t.Name()+":1111")
	t.Run("create a new client", func(t *testing.T) {
		_, err := NewSSHClient(testSSHConfig)
		if err == nil {
			t.Fatal(err)
		}
	})

}
