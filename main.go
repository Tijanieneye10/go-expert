package main

import (
	_ "embed"
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
	Error    error
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
				results <- Result{Url: url, Error: err}
				log.Fatal(err)
			}

			resp, err := http.Get(url)

			if err != nil {
				results <- Result{Url: url, Error: err}
				_ = os.Remove(filePath)
				log.Fatal(err)
			}

			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				results <- Result{Url: url, Error: fmt.Errorf("bad status code: %d", resp.StatusCode)}
			}

			size, err := io.Copy(storePath, resp.Body)
			if err != nil {
				results <- Result{Url: url}
				_ = os.Remove(filePath)
				log.Fatal(err)
				return
			}

			results <- Result{Url: url, Filename: filename, Size: size, Duration: time.Since(start)}
			fmt.Printf("%s took %s\n", url, time.Since(start))

		}(url)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

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

var counter int = 0
var mu sync.Mutex

func worker(wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	counter = counter + 1
}

//go:embed output.txt
var content string

func main() {

	fmt.Println(content)

	//
	//myFile, err := os.Open("./output.txt")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//_, err = myFile.Write([]byte("Hello World\n"))
	//
	//if err != nil {
	//	return
	//}
	//
	//var wg sync.WaitGroup
	//
	//for i := 0; i < 1000; i++ {
	//	wg.Add(1)
	//	go worker(&wg)
	//
	//}
	//
	//wg.Wait()
	//fmt.Println("the value of x ", counter)
}
