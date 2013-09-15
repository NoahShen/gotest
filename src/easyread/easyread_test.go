package easyread

import (
	"fmt"
	"testing"
)

func _TestEasyLogin(t *testing.T) {
	session := &EasyreadSession{}
	err := session.Login("username", "password")
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetSubSummary(t *testing.T) {
	session := &EasyreadSession{}
	err := session.Login("username", "password")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("userId:", session.UserId)
	fmt.Println("username:", session.Username)
	subSummary, getSummaryErr := session.GetSubSummary()
	if getSummaryErr != nil {
		t.Fatal(getSummaryErr)
	}
	fmt.Printf("subsummary: %v\n", subSummary)
	fmt.Println("============")

	newsFeed, newsSourceErr := session.GetNewsSource(subSummary.Entries[0].Id)
	if newsSourceErr != nil {
		t.Fatal(newsSourceErr)
	}
	for _, newsEntry := range newsFeed.Entries {
		fmt.Printf("newsEntry: %+v\n", newsEntry)
	}

}
