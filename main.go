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
	ErrorCode int
	ErrorMsg  string
}

func (e errorPage) Error() string {
	return e.ErrorMsg
}

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))
var artistTmpl = template.Must(template.ParseFiles("templates/artist.html"))
var errTmpl = template.Must(template.ParseFiles("templates/error.html"))

func main() {
	err := getArtistsData()
	if (err != errorPage{}) {
		http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			errTmpl.Execute(w, err)
		})
	}
	http.Handle("/static/", http.FileServer(http.Dir("assets/")))
	http.HandleFunc("/", homeHandler)

	port := "localhost:8081"
	log.Println("Listening on http://" + port)
	http.ListenAndServe(port, nil)
}
