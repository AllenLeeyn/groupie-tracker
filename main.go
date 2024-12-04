package main

import (
	"html/template"
	"log"
	"net/http"
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
	ErrorCode int
	ErrorMsg  string
}

func (e errorPage) Error() string {
	return e.ErrorMsg
}

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
