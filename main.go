package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func DownloadFile(url, path string) error {
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

	startTime := time.Now()

	for _, url := range multipleUrl {
		directoryUpload := "./upload"
		err := os.MkdirAll(directoryUpload, os.ModePerm)

		if err != nil {
			log.Fatal(err)
		}

		err = DownloadFile(url, directoryUpload)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("it took:", time.Since(startTime))
}
