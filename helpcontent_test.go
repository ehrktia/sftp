package helpcontent
import "testing"
func TestHelpcontent(t *testing T){
	filname := filepath.Join("helpcontent", "helpcontent")
	content, err := ioutil.ReadFile(Filnam)
	if err != nil {
		log.Fatal("ERROR: Reading helpcontent file")
	}
	inp:= string(content)

}
