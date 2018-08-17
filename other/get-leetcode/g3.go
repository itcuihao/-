package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func linkScrape() {
	doc, err := goquery.NewDocument("https://leetcode.com/problemset/all/")
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("td").Each(func(index int, item *goquery.Selection) {
		linkTag := item
		fmt.Println(len(linkTag.Nodes))
		link, _ := linkTag.Attr("href")
		linkText := linkTag.Text()
		fmt.Printf("Link #%d: '%s' - '%s'\n", index, linkText, link)
	})
}

func main() {
	linkScrape()
}
