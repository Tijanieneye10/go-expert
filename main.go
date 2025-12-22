package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

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
