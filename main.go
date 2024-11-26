package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

const apiURL = "https://groupietrackers.herokuapp.com/api/"

type listPage struct {
	Artists   []artistData
	NbChecked string
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
	log.Println("Listening on " + port)
	http.ListenAndServe(port, nil)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	for index, artist := range artists {
		if req.URL.Path[1:] == artist.Name &&
			req.Method == http.MethodGet {
			artistHandler(w, index)
			return
		}
	}
	if req.URL.Path != "/" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	checkErr(req.ParseForm())
	membersNb, ok := req.Form["members number"]
	var homePage *listPage
	if ok {
		if req.Method != "POST" {
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
			return
		}
		filteredArtists := filter(artistsData, membersNb[0], checkMembersNb)
		homePage = &listPage{
			Artists:   filteredArtists,
			NbChecked: membersNb[0],
		}
	} else {
		if req.Method != "GET" {
			http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
			return
		}
		homePage = &listPage{
			Artists: artistsData,
		}
	}

	indexTmpl.Execute(w, homePage)
}

func artistHandler(w http.ResponseWriter, index int) {
	artPage := &artistPage{Artist: artistsData[index]}

	artistTmpl.Execute(w, artPage)
}
