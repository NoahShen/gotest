package html

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"testing"
)

// This example scrapes the 10 reviews shown on the home page of MetalReview.com,
// the best metal review site on the web :) (and no, I'm not affiliated to them!)
func _TestExampleScrape_MetalReview(t *testing.T) {
	// Load the HTML document (in real use, the type would be *goquery.Document)
	var doc *goquery.Document
	var e error

	if doc, e = goquery.NewDocument("http://www.shjjcd.gov.cn/www/310118/"); e != nil {
		panic(e.Error())
	}

	// Find the review items (the type of the Selection would be *goquery.Selection)
	doc.Find("#WebSitePrice tr").FilterFunction(func(i int, s *goquery.Selection) bool {
		return !s.HasClass("title")
	}).Find("td:nth-child(1)").Each(func(i int, s *goquery.Selection) {
		name := s.Text()
		fmt.Printf("%s\n", name)
	})
	// To see the output of the Example while running the test suite (go test), simply
	// remove the leading "x" before Output on the next line. This will cause the
	// example to fail (all the "real" tests should pass).

	// xOutput: voluntarily fail the Example output.
}
