package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("https://leetcode.com/problemset/all/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	// #question-app > div > div:nth-child(2) > div.question-list-base > div.table-responsive.question-list-table > table > tbody.reactable-data
	// Find the review items
	doc.Find("body tr").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		ah, _ := s.Attr("href")
		a := s.Text()
		fmt.Printf("Review %d: %s %s\n", i, a, ah)
	})
}

func main() {
	ExampleScrape()
}
