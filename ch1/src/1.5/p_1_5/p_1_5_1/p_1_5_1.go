package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {

		fmt.Printf("Provided URL %s", arg)
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v", err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
