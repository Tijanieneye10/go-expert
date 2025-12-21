package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

	url := "https://images.pexels.com/photos/14733025/pexels-photo-14733025.jpeg"

	err := DownloadFile(url, "./download")

	if err != nil {
		log.Fatal(err)
	}
}
