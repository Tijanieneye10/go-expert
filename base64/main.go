package base64

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := "Hello world, what is going on here"

	resp := base64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println("Here is encoded data", resp)

	decodedData, _ := base64.StdEncoding.DecodeString(resp)

	fmt.Println("Here is the decoded data", string(decodedData))
}
