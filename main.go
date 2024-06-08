package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown"
)

func urlToMarkdown(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("error fetching URL: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	converter := htmltomarkdown.NewConverter("", true, nil)
	markdown, err := converter.ConvertString(string(body))
	if err != nil {
		return "", fmt.Errorf("error converting HTML to Markdown: %v", err)
	}

	return markdown, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL.")
		return
	}

	url := os.Args[1]
	markdownContent, err := urlToMarkdown(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	filename := "output.md"
	err = os.WriteFile(filename, []byte(markdownContent), 0644)
	if err != nil {
		fmt.Println("Error writing Markdown file:", err)
		return
	}

	fmt.Println("Markdown saved to:", filename)
}
