package main

import (
	"log"
	"net/http"
	"text/template"
)

const apiURL = "https://groupietrackers.herokuapp.com/api/"

type listPage struct {
	Artists []artistData
	SortBy  string
}

type artistPage struct {
	Artist artistData
}

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))
var artistTmpl = template.Must(template.ParseFiles("templates/artist.html"))

func main() {
	getArtist()
	http.Handle("/static/", http.FileServer(http.Dir("assets/")))

	http.HandleFunc("/", homeHandler)

	port := "localhost:8081"
	log.Println("Listening on http://" + port)
	http.ListenAndServe(port, nil)
}
