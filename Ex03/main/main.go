package main

import (
	"flag"
	"fmt"
	"os"
	"parser/parser"
)

func main() {
	fileName := flag.String("file", "index.html", "the html file to parse")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		exit(fmt.Sprintf("Error opening the csv file: %s", *fileName))
	}
	data := parser.HtmlParser(file)
	fmt.Printf("%v\n", data)
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
