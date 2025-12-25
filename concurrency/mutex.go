package concurrency

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

type Store struct {
	Url   string
	Error error
}

func main() {
	urls := []string{
		"https://img.freepik.com/free-photo/laptop-with-sun-background_1232-429.jpg",
		"https://hips.hearstapps.com/hmg-prod/images/dutch-colonial-house-style-66956274903da.jpg",
	}

	var wg sync.WaitGroup
	var mux sync.Mutex

	response := make(chan *Store)

	limiter := make(chan struct{}, 2)

	for _, url := range urls {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			limiter <- struct{}{}

			defer func() { <-limiter }()

			mux.Lock()

			defer mux.Unlock()

			filePath := filepath.Base(url)

			storeDirectory := filepath.Join("./upload", filePath)

			resp, err := http.Get(url)

			if err != nil {
				response <- &Store{Url: url, Error: err}
				log.Fatal(err)
			}

			defer resp.Body.Close()

			file, err := os.Create(storeDirectory)

			if err != nil {
				response <- &Store{Url: url, Error: err}
				log.Fatal(err)
			}

			_, err = io.Copy(file, resp.Body)

			if err != nil {
				response <- &Store{Url: url, Error: err}
				log.Fatal(err)
			}
		}(url)
	}

	wg.Wait()
}
