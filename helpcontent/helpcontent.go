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
		l.Panicf(fmt.Sprintf("%s-%v", "can not retrieve dir", err))
	}
	fileName := filepath.Join(pwd,  "helpcontent")
	l.Printf("opening file:%s\n", fileName)
	return os.Open(fileName)
}

func Helpcontent(l *log.Logger) string {
	f, err := getFile(l)
	defer func() {
		if err := f.Close(); err != nil {
			panic("error closing help file")
		}
	}()
	if err != nil {
		l.Fatalf("%s-%s\n", "ERROR", fmt.Sprintf("error reading file:%v,received error:%v\n", f, err))
	}
	d, err := io.ReadAll(f)
	if err != nil && err != io.EOF {
		l.Fatalf("%s-%s", "ERROR", err)
	}
	return string(d)
}
