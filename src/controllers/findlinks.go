package controllers

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func Parse(url string) (*html.Node, error) {
	fmt.Print(url)
	r, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Cannot get the page")
	}

	b, err := html.Parse(r.Body)

	if err != nil {
		return nil, fmt.Errorf("Can not parse the page")
	}
	return b, err
}

func PageLinks(links []string, n *html.Node) []string {

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = PageLinks(links, c)
	}
	return links
}
