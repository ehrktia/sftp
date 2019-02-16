package pwdfilgen

import (
	"encoding/hex"
	"errors"
	io "io/ioutil"
	"log"
	"os"
)

/* var declaration */

var err = errors.New("ERR Occured")

var Filnam = "pwd"

/* main operations */
func GenPwdFile(pwd []byte) {
	inpwd := pwd
	_, err := os.OpenFile(Filnam, os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("Error: generating file -", err)
	}
	encodedStr := hex.EncodeToString(inpwd)
	p := []byte(encodedStr)
	io.WriteFile(Filnam, p, 0644)

}
func DecodePwd() string {
	pwdfil, err := io.ReadFile(Filnam)

	if err != nil {
		log.Fatal("ERROR: Reading pwdfile -", err)
	}
	src := []byte(pwdfil)
	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal("ERROR: Decoding pwd from file - ", err)
	}
	return string(dst[:n])

}
