package main

import (
	"log"
	"net/http"
	"sync"
)

const (
	targetURL   = "https://iq.vntu.edu.ua"
	numRequests = 5000
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	inProgress := 0

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(requestID int) {
			defer wg.Done()

			// Increment the inProgress counter
			mu.Lock()
			inProgress++
			log.Printf("Request %d started, in progress: %d", requestID, inProgress)
			mu.Unlock()

			// Make the GET request
			resp, err := http.Get(targetURL)
			if err != nil {
				log.Printf("Request %d failed: %v", requestID, err)
			} else {
				log.Printf("Request %d completed with status code: %d", requestID, resp.StatusCode)
				resp.Body.Close()
			}

			// Decrement the inProgress counter
			mu.Lock()
			inProgress--
			log.Printf("Request %d ended, in progress: %d", requestID, inProgress)
			mu.Unlock()
		}(i)
	}

	// Wait for all requests to complete
	wg.Wait()
	log.Println("All requests completed.")
}
