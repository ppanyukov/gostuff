package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	uris := os.Args[1:]
	for _, uri := range uris {
		fmt.Println("URL:", uri)
		r, err := http.Get(uri)

		if err == nil {
			fmt.Printf("OK - %s: %s", uri, r.Status)
			bytes, err := ioutil.ReadAll(r.Body)
			if err == nil {
				fmt.Printf(" %d bytes", len(bytes))
			} else {
				fmt.Printf(" - error reading body: %s", err)
			}
		} else {
			fmt.Printf("ERR - %s: %s", uri, err)
		}

		fmt.Println()
	}
}
