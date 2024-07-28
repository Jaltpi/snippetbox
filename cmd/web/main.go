package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// Get Reqeusts
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /intro/{$}", intro)
	mux.HandleFunc("GET /snippet/view/{id}/{$}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	// Post Requests
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	hostport := "localhost:4000"
	log.Println("Starting server on :4000")
	log.Printf("Access development link by copying this output into your web browser -> %v\n",
		hostport)
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
