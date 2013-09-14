package easyread

import (
	"fmt"
	"testing"
)

func _TestEasyLogin(t *testing.T) {
	session := &EasyreadSession{}
	err := session.login("username", "password")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetSubSummary(t *testing.T) {
	session := &EasyreadSession{}
	err := session.login("username", "password")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("userId:", session.UserId)
	fmt.Println("username:", session.Username)
	session.getSubSummary()
}
