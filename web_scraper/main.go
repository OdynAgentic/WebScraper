package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <URL or search query>")
		os.Exit(1)
	}
	input := strings.Join(os.Args[1:], " ")

	// Check if the input is a URL
	_, err := url.ParseRequestURI(input)
	var targetURL string
	if err != nil {
		// If it's not a URL, treat it as a search query for Google
		targetURL = "https://www.google.com/search?q=" + url.QueryEscape(input)
	} else {
		targetURL = input
	}

	// create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// create chrome instance
	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var text string
	err = chromedp.Run(ctx,
		chromedp.Navigate(targetURL),
		// Wait for the main content to load
		chromedp.WaitVisible("body", chromedp.ByQuery),
		// Extract text from the main content
		chromedp.Text("body", &text, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Clean up the text
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "\n\n\n", "\n")
	text = strings.ReplaceAll(text, "\n\n", "\n")

	fmt.Println(text)
}