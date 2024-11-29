package main

import (
	"net/http"
	"sort"
)

// homeHandler() checks the http request methods and url
// and provide the correct response.
func homeHandler(w http.ResponseWriter, req *http.Request) {
	for index, artist := range artistsLst {
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
	sortLst(sortCriteria)

	// POST method is used for sorting list.
	// Invalid request is ignore and use default settings.
	if req.Method == http.MethodPost {
		sortCriteria = req.FormValue("sort")
		if sortCriteria != "creation_date" && sortCriteria != "name" {
			sortCriteria = "default"
		}
		sortLst(sortCriteria)
		pageOrder := req.FormValue("switch-order")
		if pageOrder == "▼" {
			order = "▲"
			revLst()
		} else if pageOrder == "▲" {
			order = "▼"
		}
	} else if req.Method != http.MethodGet {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	homePage := &listPage{
		Artists: artistsLst,
		SortBy:  sortCriteria,
		Order:   order}
	indexTmpl.Execute(w, homePage)
}

// artistHandler() generates the html response for an artist
func artistHandler(w http.ResponseWriter, index int) {
	artPage := &artistPage{Artist: artistsLst[index]}
	artistTmpl.Execute(w, artPage)
}

// sortLst() sorts artistLst based on sortCriteria in ascending order.
func sortLst(sortCriteria string) {
	sort.Slice(artistsLst, func(i, j int) bool {
		switch sortCriteria {
		case "name":
			return artistsLst[i].Name < artistsLst[j].Name
		case "creation_date":
			return artistsLst[i].CreationDate < artistsLst[j].CreationDate
		}
		return artistsLst[i].Id < artistsLst[j].Id
	})
}

// revLst() reverses artistsLst
func revLst() {
	start, end := 0, len(artistsLst)-1
	for start < end {
		artistsLst[start], artistsLst[end] =
			artistsLst[end], artistsLst[start]
		start++
		end--
	}
}
