package gom

import "testing"

func TestHello(t *testing.T) {
	t.Log("into test hello")
	result := Hello()
	if result != true {
		t.Errorf(" invalid")
	}
}
