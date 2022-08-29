package travisTest

import "testing"

func Testhellogolang(t *testing.T)  {
	if hellogolang()=="Hello,World!" {
		t.log("Test Passed")
	}else{
		t.Error("Test Failed")
	}
}