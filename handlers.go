package main

import (
	"net/http"
	"sort"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	for index, artist := range artistsData {
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
	sortCriteria := req.URL.Query().Get("sort")
	if sortCriteria != "creation_date" {
		sortCriteria = "name"
	}
	sortList(artistsData, sortCriteria)

	homePage := &listPage{Artists: artistsData, SortBy: sortCriteria}

	indexTmpl.Execute(w, homePage)
}

func artistHandler(w http.ResponseWriter, index int) {
	artPage := &artistPage{Artist: artistsData[index]}

	artistTmpl.Execute(w, artPage)
}

func sortList(artistsData []artistData, sortCriteria string) {
	sort.Slice(artistsData, func(i, j int) bool {
		if sortCriteria == "name" {
			return artistsData[i].Name < artistsData[j].Name
		}
		return artistsData[i].CreationDate < artistsData[j].CreationDate
	})
}
