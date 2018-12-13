package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/daetal-us/kii"
)

func main() {
	flag.Parse()
	url := flag.Arg(0)
	if url == "" {
		fmt.Println("A URL is required.")
		os.Exit(1)
	}
	results, err := kii.FromURL(url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	encoded, err := json.Marshal(results)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(encoded))
}
