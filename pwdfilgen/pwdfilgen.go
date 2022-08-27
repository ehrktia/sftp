package pwdfilgen

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

var Filnam = "pwd"

func GenPwdFile(pwd []byte) error{
	inpwd := pwd
	_, err := os.OpenFile(Filnam, os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal("Error: generating file -", err)
	}
	encodedStr := hex.EncodeToString(inpwd)
	p := []byte(encodedStr)
return	os.WriteFile(Filnam,p,0644)

}
func DecodePwd() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("error getting current working dir:%v\n", err))
	}
	fpath := filepath.Join(cwd, "pwdfilgen", "pwd")
	f, err := os.Open(fpath)
	if err != nil {
		panic(fmt.Sprintf("can not read file:%v\n", fpath))
	}
	pwdfil, err := io.ReadAll(f)
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
