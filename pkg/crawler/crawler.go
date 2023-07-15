package crawler

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

func Crawler(urlChannel chan string, visitedURLs map[string]bool, visitedURLsMutex *sync.Mutex, done chan bool) {
	for url := range urlChannel {
		visitedURLsMutex.Lock()
		if visitedURLs[url] {
			continue
		}

		visitedURLs[url] = true

		resp, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer resp.Body.Close()

		doc, err := html.Parse(resp.Body)
		if err != nil {
			fmt.Println(err)
			continue
		}
		links := extractLinks(doc)
		for _, link := range links {
			fmt.Println(link)
			go func(link string) {
				urlChannel <- link
			}(link)
		}
	}
	done <- true
}

func extractLinks(n *html.Node) []string {
	var links []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				link := strings.TrimSpace(attr.Val)
				if link != "" {
					links = append(links, link)
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, extractLinks(c)...)
	}
	return links
}
