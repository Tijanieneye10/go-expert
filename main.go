package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	url := "https://images.pexels.com/photos/14733025/pexels-photo-14733025.jpeg"

	filePath := filepath.Base(url)

	directory := filepath.Join("./", filePath)

	storeHere, err := os.Create(directory)

	if err != nil {
		panic(err)
	}

	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	defer response.Body.Close()

	_, err = io.Copy(storeHere, response.Body)

	if err != nil {
		return
	}

}
