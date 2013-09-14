package html

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"testing"
)

var allCount = 0

func TestGetFoodNutrition(t *testing.T) {
	// Load the HTML document (in real use, the type would be *goquery.Document)
	var doc *goquery.Document
	var e error
	host := "http://yingyang.911cha.com/"
	if doc, e = goquery.NewDocument(host); e != nil {
		panic(e.Error())
	}

	// Find the review items (the type of the Selection would be *goquery.Selection)
	doc.Find("strong.f14").FilterFunction(func(i int, s *goquery.Selection) bool {
		name := s.Text()
		return name == "食物分类"
	}).SiblingsFiltered("a.f14").Each(func(i int, s *goquery.Selection) {
		//typeName := s.Text()
		href, _ := s.Attr("href")
		//fmt.Printf("typeName: %s url: %s\n", typeName, href)
		url := host + href
		getFood(url)
	})
	fmt.Printf("all count: %d\n", allCount)
}

func getFood(url string) {
	// Load the HTML document (in real use, the type would be *goquery.Document)
	var doc *goquery.Document
	var e error
	if doc, e = goquery.NewDocument(url); e != nil {
		panic(e.Error())
	}

	// Find the review items (the type of the Selection would be *goquery.Selection)
	doc.Find("div.mtitle").FilterFunction(func(i int, s *goquery.Selection) bool {
		a := s.Find("a")
		name := a.Text()
		return name == "食物营养成分查询"
	}).SiblingsFiltered("div.mcon").Find("a").Each(func(i int, s *goquery.Selection) {
		foodName := s.Text()
		href, _ := s.Attr("href")
		fmt.Printf("foodName: %s url: %s\n", foodName, href)
		allCount++
	})
}
