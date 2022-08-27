package helpcontent

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func getFile(l *log.Logger) (*os.File, error) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("%s-%v", "can not retrieve dir", err))
	}
	fileName := filepath.Join(pwd, "helpcontent", "helpcontent")
	return os.Open(fileName)
}

func Helpcontent(l *log.Logger) string {
	f, err := getFile(l)
	if err != nil {
		l.Printf("%s-%s\n", "ERROR", fmt.Sprintf("error reading file:%v,received error:%v\n", f, err))
	}
	d, err := io.ReadAll(f)
	if err != nil && err != io.EOF {
		l.Printf("%s-%s", "ERROR", err)
	}
	return string(d)
}
