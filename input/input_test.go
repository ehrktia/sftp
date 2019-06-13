package input

import "testing"
func TestSetUserInp(t *testing.T) {
	inp := []struct {
		uname  string
		pwd    string
		url    string
		srcdir string
		tgtdir string
	}
	{
		{"tester","test!ing","testing.com","/src","/tgt"},
		{"tester1","test!ing1","testing1.com","/src1","/tgt1"}
	}
	for _,val:=range inp{
		
	}

}
