package input

type inpparms struct {
	Uname  string
	Pwd    string
	Url    string
	SrcDir string
	TgtDir string
}

var Inp = new(inpparms)

func SetUserInp(uname, pwd, url, srcdir, tgtdir string) {
	(*Inp).Uname = uname
	(*Inp).Pwd = pwd
	(*Inp).Url = url
	(*Inp).SrcDir = srcdir
	(*Inp).TgtDir = tgtdir
}
