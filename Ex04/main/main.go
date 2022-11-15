package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"os"
	"scrapper/scrapper"
)

type locs struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []locs `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	url := flag.String("url", "https://gophercises.com/", "url to build a sitemap of")
	depth := flag.Int("depth", 4, "maximum depth to go")
	flag.Parse()
	data := scrapper.Scrapper(*url, *depth)
	fmt.Print(xml.Header)
	toxml := urlset{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
	}
	for _, d := range data {
		toxml.Urls = append(toxml.Urls, locs{d.HREF})
	}
	xmlen := xml.NewEncoder(os.Stdout)
	xmlen.Indent("", "  ")
	if err := xmlen.Encode(toxml); err != nil {
		panic(err)
	}
	fmt.Println()
}
