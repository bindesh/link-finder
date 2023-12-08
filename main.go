package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bindesh/link-finder/src/controllers"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("q")
	fmt.Fprintf(w, "Page=%q\n", url)
	if len(url) == 0 {
		return
	}
	page, err := controllers.Parse("https://" + url)

	if err != nil {
		fmt.Printf("Error getting the page %s %s\n", url, err)
	}

	links := controllers.PageLinks(nil, page)
	for _, links := range links {
		fmt.Fprintf(w, "Link = %q\n", links)
	}
}
