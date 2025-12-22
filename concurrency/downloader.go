package concurrency

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

func DownloadFile(url, path string, wg *sync.WaitGroup) error {

	defer wg.Done()
	filename := filepath.Base(url)

	storePath := filepath.Join(path, filename)

	resp, err := http.Get(url)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	output, err := os.Create(storePath)

	if err != nil {
		return err
	}

	defer output.Close()

	_, err = io.Copy(output, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	//url := "https://images.pexels.com/photos/14733025/pexels-photo-14733025.jpeg"
	//
	//err := DownloadFile(url, "./download")

	multipleUrl := []string{
		"https://images.pexels.com/photos/14733025/pexels-photo-14733025.jpeg",
		"https://hips.hearstapps.com/hmg-prod/images/dutch-colonial-house-style-66956ed4ce458.jpg",
		"https://hips.hearstapps.com/hmg-prod/images/edc100123egan-002-6500742f5feb7.jpg",
	}

	var wg sync.WaitGroup

	startTime := time.Now()

	for _, url := range multipleUrl {
		wg.Add(1)
		directoryUpload := "./upload"
		err := os.MkdirAll(directoryUpload, os.ModePerm)

		if err != nil {
			log.Fatal(err)
		}

		go func() {
			err = DownloadFile(url, directoryUpload, &wg)
			if err != nil {
				log.Fatal(err)
			}
		}()

	}

	wg.Wait()

	fmt.Println("it took:", time.Since(startTime))
}
