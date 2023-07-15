package main

import (
	"search-engine/pkg/crawler"
	"sync"
)

func main() {
	startURL := "https://www.google.com"

	// Crear un mapa para almacenar las URLs visitadas
	visitedURLs := make(map[string]bool)
	visitedURLsMutex := sync.Mutex{}
	// Crear un canal para comunicarse entre las goroutines
	urlChannel := make(chan string)
	done := make(chan bool)
	// Crear un grupo de espera para sincronizar las goroutines
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		crawler.Crawler(urlChannel, visitedURLs, &visitedURLsMutex, done)
	}()

	urlChannel <- startURL

	wg.Wait()
	close(urlChannel)
	<-done
}
