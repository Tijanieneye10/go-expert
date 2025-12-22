package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Result struct {
	Url      string
	Filename string
	Size     int64
	Duration time.Duration
}

func ConcurrentDownloader(urls []string, destPath string, maxConcurrent int) {
	if err := os.MkdirAll(filepath.Dir(destPath), 0777); err != nil {
		log.Fatal(err)
	}

	results := make(chan Result)

	var wg sync.WaitGroup

	limiter := make(chan struct{}, maxConcurrent)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			limiter <- struct{}{}

			defer func() { <-limiter }()

			start := time.Now()

			filename := filepath.Base(url)
			filePath := filepath.Join(destPath, filename)

			storePath, err := os.Create(filePath)

			if err != nil {
				log.Fatal(err)
			}

			resp, err := http.Get(url)

			if err != nil {
				_ = os.Remove(filePath)
				log.Fatal(err)
			}

			defer resp.Body.Close()

			_, err = io.Copy(storePath, resp.Body)
			if err != nil {
				_ = os.Remove(filePath)
				log.Fatal(err)
				return
			}

			fmt.Printf("%s took %s\n", url, time.Since(start))

		}(url)
	}

}

func DownloadFile(url, path string, wg *sync.WaitGroup) error {

	defer wg.Done()
	filename := filepath.Base(url)

	storePath := filepath.Join(path, filename)

	resp, err := http.Get(url)

	if err != nil {
		_ = os.Remove(storePath)
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		_ = os.Remove(storePath)
	}

	output, err := os.Create(storePath)

	if err != nil {
		_ = os.Remove(storePath)
		return err
	}

	defer output.Close()

	_, err = io.Copy(output, resp.Body)
	if err != nil {
		_ = os.Remove(storePath)
		return err
	}

	return nil
}

func main() {

}
