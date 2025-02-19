package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/gocolly/colly/v2"
)

// create a struct where we can store the scraped data
type Article struct {
	URL     string `json:"url"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

func main() {
	//start timer
	start := time.Now()

	// List of wikipedia urls about AI companies
	urls := []string{
		"https://en.wikipedia.org/wiki/Nvidia",
		"https://en.wikipedia.org/wiki/OpenAI",
		"https://en.wikipedia.org/wiki/Anthropic",
		"https://en.wikipedia.org/wiki/DeepMind",
		"https://en.wikipedia.org/wiki/XAI_(company)",
		"https://en.wikipedia.org/wiki/DeepSeek",
		"https://en.wikipedia.org/wiki/Baidu",
		"https://en.wikipedia.org/wiki/Microsoft",
		"https://en.wikipedia.org/wiki/Alphabet_Inc.",
		"https://en.wikipedia.org/wiki/Amazon_(company)",
	}

	// use slize to store the scraped articles
	var articles []Article
	// had issue with writes to the slice at the same time so found this solution
	var mu sync.Mutex

	// Start a colley and ensure with is set to async and allows wikipedia
	c := colly.NewCollector(
		colly.Async(true),
		colly.AllowedDomains("en.wikipedia.org"),
	)

	// use a call back to start the process by creating a new article for the url being visited
	c.OnRequest(func(r *colly.Request) {
		article := &Article{URL: r.URL.String()}
		r.Ctx.Put("article", article)
		fmt.Println("Visiting:", r.URL)
	})

	// Call the url and find the title of the article
	c.OnHTML("h1#firstHeading", func(e *colly.HTMLElement) {
		if art, ok := e.Request.Ctx.GetAny("article").(*Article); ok {
			art.Title = e.Text
		}
	})

	c.OnHTML("div#mw-content-text div.mw-parser-output p", func(e *colly.HTMLElement) {
		// Append each paragraph to the Article's content.
		if art, ok := e.Request.Ctx.GetAny("article").(*Article); ok {
			art.Content += e.Text + "\n"
		}
	})

	// once everything has been scrappend append the article to the slice and print the finished URL
	c.OnScraped(func(r *colly.Response) {
		if art, ok := r.Ctx.GetAny("article").(*Article); ok {
			mu.Lock()
			articles = append(articles, *art)
			mu.Unlock()
			fmt.Println("Finished:", art.URL)
		}
	})

	// create errory logging for colly
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("Error on %s: %v", r.Request.URL, err)
	})

	// Start the crawling useing all the steps above and iterate through the urls
	for _, url := range urls {
		c.Visit(url)
	}

	// wait for all requests to complete
	c.Wait()

	// create JSON file
	outFile, err := os.Create("output.jsonl")
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer outFile.Close()

	// write the articles to the JSON file
	encoder := json.NewEncoder(outFile)
	for _, art := range articles {
		if err := encoder.Encode(art); err != nil {
			log.Printf("Error. URL: %s: %v", art.URL, err)
		}
	}

	fmt.Println("Complete. Review at:  output.jsonl")
	elapsed := time.Since(start)
	fmt.Printf("Timer To Run: %s\n", elapsed)
}
