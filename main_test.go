package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gocolly/colly/v2"
)

func TestAnthropic(t *testing.T) {
	//set vars for testing
	url := "https://en.wikipedia.org/wiki/Anthropic"
	expectedTitle := "Anthropic"

	var article Article

	// Creat new colly and allow wikipedia
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)
	//create a new article for the url
	c.OnRequest(func(r *colly.Request) {
		article.URL = r.URL.String()
	})

	//Get title from the HTML
	c.OnHTML("h1#firstHeading", func(e *colly.HTMLElement) {
		article.Title = strings.TrimSpace(e.Text)
	})

	//Scrape URL and run callbacks
	if err := c.Visit(url); err != nil {
		t.Fatalf("Failed to visit URL: %v", err)
	}
	c.Wait()

	//Print out test and create error if they do not match
	fmt.Print("expected title: ", expectedTitle, "\n")
	fmt.Print("actual title: ", article.Title, "\n")
	if article.Title != expectedTitle {
		t.Errorf("Expected title %q, but got %q", expectedTitle, article.Title)
	}
}
