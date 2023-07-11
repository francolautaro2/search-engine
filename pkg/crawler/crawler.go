package crawler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/net/html"
)

func Downloader(url string, filename string) (string, error) {

	var client http.Client
	fmt.Println("Downloading...", url)
	resp, err := client.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", err
	}
	return filename, nil
}

func Crawl(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}

	links := findLinks(doc)
	for _, link := range links {
		time.Sleep(1 * time.Millisecond)
		fmt.Println("link found ->", link)
	}
}

func findLinks(n *html.Node) []string {
	var NewLinks []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				NewLinks = append(NewLinks, attr.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		NewLinks = append(NewLinks, findLinks(c)...)
	}
	return NewLinks
}
