package main

import (
	"fmt"
	"search-engine/pkg/crawler"
	"search-engine/pkg/utils"
)

func main() {
	fmt.Println("Crawler is running")
	urls, _ := utils.ReadTxtUrl("urls.txt")
	for _, u := range urls {
		file, err := utils.CreateHtmlFile(u)
		if err != nil {
			fmt.Println(err)
		}
		f, err := crawler.Downloader(u, file)
		if err != nil {
			fmt.Println(err)
		}
		crawler.Crawl(f)
	}

}
