package main

import (
	"html/template"
	"log"
	"net/http"
)

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))
var artistTmpl = template.Must(template.ParseFiles("templates/artist.html"))
var errTmpl = template.Must(template.ParseFiles("templates/error.html"))
var ArtistErr = getArtistsData()

func main() {
	http.Handle("/static/", http.FileServer(http.Dir("assets/")))
	http.HandleFunc("/", homeHandler)

	port := "10000"
	log.Println("Listening on http://localhost:" + port)
	http.ListenAndServe(":"+port, nil)
}
