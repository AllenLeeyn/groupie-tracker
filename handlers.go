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

	order := "▼"
	sortCriteria := "default"
	if req.Method == http.MethodPost {
		sortCriteria = req.FormValue("sort")

		if sortCriteria != "creation_date" && sortCriteria != "name" {
			sortCriteria = "default"
		}
		sortList(artistsData, sortCriteria)
		pageOrder := req.FormValue("switch-order")
		if pageOrder == "▼" {
			order = "▲"
			revList(artistsData)
		} else if pageOrder == "▲" {
			order = "▼"
		}

	} else if req.Method != "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	homePage := &listPage{Artists: artistsData, SortBy: sortCriteria, Order: order}
	indexTmpl.Execute(w, homePage)
}

func artistHandler(w http.ResponseWriter, index int) {
	artPage := &artistPage{Artist: artistsData[index]}
	artistTmpl.Execute(w, artPage)
}

func sortList(artistsData []artistData, sortCriteria string) {
	sort.Slice(artistsData, func(i, j int) bool {
		switch sortCriteria {
		case "name":
			return artistsData[i].Name < artistsData[j].Name
		case "creation_date":
			return artistsData[i].CreationDate < artistsData[j].CreationDate
		case "default":
			return artistsData[i].Index < artistsData[j].Index
		}
		return artistsData[i].CreationDate < artistsData[j].CreationDate
	})
}

func revList(artistData []artistData) {
	start, end := 0, len(artistData)-1

	for start < end {
		artistData[start], artistData[end] = artistData[end], artistData[start]
		start++
		end--
	}
}
