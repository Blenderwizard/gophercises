package scrapper

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	HREF string
	Text string
}

func extractText(n *html.Node) string {
	var ret string
	if n.Type == html.TextNode && n.Data != "a" {
		return strings.TrimSpace(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = fmt.Sprintf("%s %s", strings.TrimSpace(ret), extractText(c))
	}
	return strings.TrimSpace(ret)
}

func HtmlParser(r io.Reader) []Link {
	doc, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	var f func(*html.Node)
	var ar []Link
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
		textLoop:
			for _, data := range n.Attr {
				if data.Key == "href" {
					ar = append(ar, Link{
						HREF: data.Val,
						Text: extractText(n),
					})
					break textLoop
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return ar
}
