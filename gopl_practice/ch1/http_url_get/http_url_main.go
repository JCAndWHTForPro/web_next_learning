package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	inputUrlStrings := os.Args[1:]
	for _, url := range inputUrlStrings {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error:%v\n", err)
			os.Exit(1)
		}
		all, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch error:%v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%s", all)
	}
}
