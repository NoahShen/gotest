package easyread

import (
	"code.google.com/p/go.net/html"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
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

	newsFeed, newsSourceErr := session.GetNewsSource(subSummary.Entries[1].Id)
	if newsSourceErr != nil {
		t.Fatal(newsSourceErr)
	}
	for _, newsEntry := range newsFeed.Entries {
		fmt.Printf("newsEntry: %+v\n", newsEntry)
	}

	fmt.Println("============")

	articleFeed, newsSourceErr := session.GetArticle(newsFeed.Entries[1].Id)
	fmt.Printf("articleFeed: %+v\n", articleFeed)

	root, htmlParseErr := html.Parse(strings.NewReader(articleFeed.Entry.Content.Content))
	if htmlParseErr != nil {
		t.Fatal(htmlParseErr)
	}

	fmt.Println("============")

	doc := goquery.NewDocumentFromNode(root)
	contentObj := doc.Find("div.fs-content")
	content := contentObj.Text()
	content = strings.Replace(content, " ", "", -1)
	fmt.Println("doc text:", content)
}
