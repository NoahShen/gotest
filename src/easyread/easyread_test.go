package easyread

import (
	"testing"
)

func TestEasyLogin(t *testing.T) {
	session := &EasyreadSession{}
	err := session.login("username", "password")
	if err != nil {
		t.Fatal(err)
	}
}
