package input

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSetUserInp(t *testing.T) {
	filname := filepath.Join("testdata", "input")
	file, err := os.Open(filname)
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content := scanner.Text()
		cols := strings.Split(content, ",")
		SetUserInp(cols[0], cols[1], cols[2], cols[3], cols[4])
	}
}
