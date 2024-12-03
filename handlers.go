package main

import (
	"net/http"
)

// homeHandler() handles all request and url.
// passes '/artistName' to artistHandler().
func homeHandler(w http.ResponseWriter, req *http.Request) {
	// fmt.Println("ARTISTSDATA:", artistsData)
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
	homePage = &listPage{
		Artists: artistsLst,
		Order:   order,
		SortBy:  sortCriteria,
	}

	indexTmpl.Execute(w, homePage)
}

// artistHandler() generates the html response for an artist
func artistHandler(w http.ResponseWriter, index int) {
	artPage := &artistPage{Artist: artistsLst[index]}
	artistTmpl.Execute(w, artPage)
}
