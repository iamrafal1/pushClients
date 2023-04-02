package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/example1", handlerWrapper("templates/example1.html"))
	http.HandleFunc("/example2", handlerWrapper("templates/example2.html"))
	log.Fatal(http.ListenAndServe(":7070", nil))
}

func handlerWrapper(htmlFile string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		// Read in the template with SSE JavaScript code.
		t, err := template.ParseFiles(htmlFile)
		if err != nil {
			log.Fatal("error parsing template.")

		}

		// Render the template, writing to `w`.
		t.Execute(w, nil)

		// Done.
		log.Println("Finished HTTP request at", r.URL.Path)
	}
}
