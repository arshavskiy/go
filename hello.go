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
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.ListenAndServe(":80", nil)
}
