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
		http.HandleFunc("/", func (w http.ResponseWriter, req *http.Request)  {
			errTmpl.Execute(w, err)
		})
	}
	http.Handle("/static/", http.FileServer(http.Dir("assets/")))
	http.HandleFunc("/", homeHandler)

	port := "localhost:8081"
	log.Println("Listening on http://" + port)
	http.ListenAndServe(port, nil)
}

func sortArtists(req *http.Request, arr []artist) (string, string, error) {
	order := "▼"
	sortCriteria := "default"
	sortLst(arr, sortCriteria)

	// POST method is used for sorting list.
	// Invalid request is ignore and use default settings.
	if req.Method == http.MethodPost {
		sortCriteria = req.FormValue("sort")
		if sortCriteria != "creation_date" && sortCriteria != "name" {
			sortCriteria = "default"
		}
		sortLst(arr, sortCriteria)
		pageOrder := req.FormValue("switch-order")
		if pageOrder == "▼" {
			order = "▲"
			revLst(arr)
		} else if pageOrder == "▲" {
			order = "▼"
		}
	} else if req.Method != http.MethodGet {
		return "", "", MethodNotAllowedErr // should change it
	}
	return order, sortCriteria, nil
}
