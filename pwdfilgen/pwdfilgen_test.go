package pwdfilgen

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestGenPwdFile(t *testing.T) {
	//check file
	Filnam := "pwd"
	inptxt := []byte("pwd")
	GenPwdFile(inptxt) //nolint:errcheck
	_, err := os.Stat(Filnam)
	if err != nil {
		log.Fatal(err)
	}
	//check if inp txt is encoded
	encodedout := hex.EncodeToString(inptxt)
	content, err := ioutil.ReadFile(Filnam)
	if err != nil {
		log.Fatal(err)
	}
	actualfilcontent := []byte(content)
	actual := string(actualfilcontent)
	if encodedout != actual {
		t.Errorf("expected: %s received: %s", encodedout, actual)
	}

}
