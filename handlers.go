package main

import (
	"errors"
	"net/http"
	"slices"
	"sort"
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

	err := getSortedArtists(w, req)
	if err != nil {
		errPage := err.(errorPage) // type assertion
		http.Error(w, errPage.errorMsg, errPage.errorCode)
		return
	}
	indexTmpl.Execute(w, homePage)
}

// artistHandler() generates the html response for an artist
func artistHandler(w http.ResponseWriter, index int) {
	artPage := &artistPage{Artist: artistsLst[index]}
	artistTmpl.Execute(w, artPage)
}

// getSortedArtists() checks for method error and sorts/filter artists
func getSortedArtists(w http.ResponseWriter, req *http.Request) error {
	// var homePage *listPage
	var newArtistsLst []artist = slices.Clone(artistsLst)

	checkErr(req.ParseForm())
	order, sortCriteria, err := sortArtists(w, req, newArtistsLst)
	checkErr(err)
	homePage = &listPage{
		Artists: newArtistsLst,
		Order:   order,
		SortBy:  sortCriteria,
	}
	return nil
}
func sortArtists(w http.ResponseWriter, req *http.Request, arr []artist) (string, string, error) {
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
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return "", "", errors.New("http error")
	}
	return order, sortCriteria, nil
}

// sortLst() sorts artistLst based on sortCriteria in ascending order.
func sortLst(arr []artist, sortCriteria string) {
	sort.Slice(arr, func(i, j int) bool {
		switch sortCriteria {
		case "name":
			return arr[i].Name < arr[j].Name
		case "creation_date":
			return arr[i].CreationDate < arr[j].CreationDate
		}
		return arr[i].Id < arr[j].Id
	})
}

// revLst() reverses artistsLst
func revLst(arr []artist) {
	start, end := 0, len(arr)-1
	for start < end {
		arr[start], arr[end] =
			arr[end], arr[start]
		start++
		end--
	}
}
