package main

import (
	"flag"
	"fmt"
	"handler/handler"
	"net/http"
	"os"
)

func blank(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "")
}

func defaultMuxHandler() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", blank)
	return mux
}

func exit(msg string) {
	fmt.Printf("%s", msg)
	os.Exit(1)
}

func main() {
	mux := defaultMuxHandler()
	port := flag.Int("port", 8080, "port to use to host the server")
	jsonFile := flag.String("json", "gopher.json", "define what json file to use.")
	flag.Parse()

	file, err := os.Open(*jsonFile)
	if err != nil {
		exit(fmt.Sprintf("Error opening the csv file: %s", *jsonFile))
	}
	jsonHandler, err := handler.JSONHandler(file, mux)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Starting the server on :%d\n", *port)
	http.ListenAndServe(fmt.Sprintf(":%d", *port), jsonHandler)
}
