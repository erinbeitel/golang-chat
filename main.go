package main

import (
	"flag"
	"ftm"
	"log"
	"net/http"
)

func main() {
	// CL flags
	port := flag.Int("port", 8000, "port to server on")
	dir := flag.String("directory", "web/", "directory of web files")
	flag.Parse()

	// handle all requests by serving a file of the same name.
	fs := http.Dir(*dir)
	fileHandler := http.FileServer(fs)
	http.Handle("/", fileHandler)

	log.Printf("Running on port %d\n", *port)

	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	// this call blocks -- the program runs here forever
	err := http.ListenAndServe(addr, nil)
	fmt.Println(err.Error())
}
