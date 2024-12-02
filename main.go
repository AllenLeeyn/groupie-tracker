package main

import (
	"log"
	"net/http"
	"text/template"
)

const apiURL = "https://groupietrackers.herokuapp.com/api/"

type listPage struct {
	Artists []artist
	SortBy  string
	Order   string
}

var homePage *listPage = &listPage{}

type artistPage struct {
	Artist artist
}

type errorPage struct {
	errorCode int
	errorMsg  string
}

func (e errorPage) Error() string {
	return e.errorMsg
}

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))
var artistTmpl = template.Must(template.ParseFiles("templates/artist.html"))

func main() {
	getArtistsData()
	http.Handle("/static/", http.FileServer(http.Dir("assets/")))
	http.HandleFunc("/", homeHandler)

	port := "localhost:8081"
	log.Println("Listening on http://" + port)
	http.ListenAndServe(port, nil)
}
