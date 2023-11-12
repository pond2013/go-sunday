package libs

import "testing"

func TestLogin(t *testing.T) {
	Greeting()
	if Login("admin", "1234") != true {
		t.Errorf("error test fail")
	}
}
