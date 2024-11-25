package main

import (
	"log"
	"net/http"
	"sort"
	"text/template"
)

const apiURL = "https://groupietrackers.herokuapp.com/api/"

type listPage struct {
	Artists []artistData
}

type artistPage struct {
	Artist artistData
}

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))
var artistTmpl = template.Must(template.ParseFiles("templates/artist.html"))

func main() {
	getArtist()
	sort.Slice(artistsData, func (i, j int) bool {
		return artistsData[i].Name <= artistsData[j].Name
	})
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

	if req.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	homePage := &listPage{Artists: artistsData}

	indexTmpl.Execute(w, homePage)
}

func artistHandler(w http.ResponseWriter, index int) {
	artPage := &artistPage{Artist: artistsData[index]}

	artistTmpl.Execute(w, artPage)
}
