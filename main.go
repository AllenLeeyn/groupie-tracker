package main

import (
	"html/template"
	"log"
	"net/http"
)

var indexTmpl, err1 = template.ParseFiles("templates/index.html")
var artistTmpl, artistTmplErr = template.ParseFiles("templates/artist.html")
var errTmpl, err2 = template.ParseFiles("templates/error.html")
var ArtistErr *errorPage = nil

func main() {
	ArtistErr = getArtistsData()
	http.Handle("/static/", http.FileServer(http.Dir("assets/")))
	http.HandleFunc("/", homeHandler)

	port := ":8081"
	log.Println("Listening on http://localhost" + port)
	http.ListenAndServe(port, nil)
}
