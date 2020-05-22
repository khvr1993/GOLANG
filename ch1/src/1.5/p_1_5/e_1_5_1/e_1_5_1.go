package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	out := os.Stdout
	var url string
	for _, arg := range os.Args[1:] {

		if !(strings.HasPrefix(arg, "http")) {
			url = "http://" + arg
			fmt.Println("URL does not have prefix")
		} else {
			fmt.Println("URL does have prefix")
		}
		fmt.Println("Provided URL " + url + "\n")

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v", err)
			os.Exit(1)
		}
		io.Copy(out, resp.Body)
		resp.Body.Close()
		fmt.Printf("The status is %s\tand code is%d", resp.Status, resp.StatusCode)
		fmt.Println(out)
	}
}
