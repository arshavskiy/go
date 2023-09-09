package main

import (
	"fmt"
	"net/http"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	fmt.Println(quote.Opt())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":80", nil)
}
