# Week 7 Assignment: Crawling and Scraping the Web

This project is a simple Go-based web crawler and scraper. It is built as part of an assignment 7 of MSDS 431 to collect research data for an online library focused on AI Companies.

## Overview

A technology firm wants to create an online knowledge base. The idea is to collect text from various Wikipedia pages from the AI industry. Instead of using Python and Scrapy (which can be slow), this project uses Go and the Colly framework to scrape pages concurrently.

## What the Project Does

- **Crawls Wikipedia Pages:** It visits a list of Wikipedia pages about companies and research related to the AI Industry
- **Scrapes the Content:** It extracts the title and main text from each page while removing the HTML tags
- **Saves the Data:** The extracted information is saved as JSON objects in a file called `output.jsonl`

## Key Call Outs
- The scraper only processes text and ignores images.
- It is configured to work only with pages from en.wikipedia.org.
- Make sure you have an active internet connection to run the scraper.

## Running the Program
Run the project using the following command: go run main.go

When you run the program, it will:
- Visit a list of Wikipedia pages
- Scrape the title and text content from each page
- Save the scraped data to a file named output.jsonl
- After the program finishes, you can open output.jsonl to review the scraped data

## Results
The program has been tested on 10 Wikipedia pages, and the scraping process takes around 200–370 milliseconds per run. Here are specific times:

Test 1: 270.975125ms
Test 2: 237.056416ms
Test 3: 241.687834ms
Test 4: 369.283084ms
Test 5: 277.788334ms
Test 6: 286.757291ms
Test 7: 254.672417ms
Test 8: 215.789417ms
Test 9: 248.196625ms
Test 10: 261.757916ms