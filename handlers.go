package main

import (
	"fmt"
	"net/http"
	"sort"
)

// // homeHandler() checks the http request methods and url
// // and provide the correct response.
// func homeHandler(w http.ResponseWriter, req *http.Request) {
// 	for index, artist := range artistsLst {
// 		if req.URL.Path[1:] == artist.Name &&
// 			req.Method == http.MethodGet {
// 			artistHandler(w, index)
// 			return
// 		}
// 	}
// 	if req.URL.Path != "/" {
// 		http.Error(w, "404 not found", http.StatusNotFound)
// 		return
// 	}
// 	order := "▼"
// 	sortCriteria := "default"
// 	sortLst(sortCriteria)

// 	// POST method is used for sorting list.
// 	// Invalid request is ignore and use default settings.
// 	if req.Method == http.MethodPost {
// 		sortCriteria = req.FormValue("sort")
// 		if sortCriteria != "creation_date" && sortCriteria != "name" {
// 			sortCriteria = "default"
// 		}
// 		sortLst(sortCriteria)
// 		pageOrder := req.FormValue("switch-order")
// 		if pageOrder == "▼" {
// 			order = "▲"
// 			revLst()
// 		} else if pageOrder == "▲" {
// 			order = "▼"
// 		}
// 	} else if req.Method != http.MethodGet {
// 		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	homePage := &listPage{
// 		Artists: artistsLst,
// 		SortBy:  sortCriteria,
// 		Order:   order}
// 	indexTmpl.Execute(w, homePage)
// }

// artistHandler() generates the html response for an artist
func artistHandler(w http.ResponseWriter, index int) {
	artPage := &artistPage{Artist: artistsLst[index]}
	artistTmpl.Execute(w, artPage)
}

// sortLst() sorts artistLst based on sortCriteria in ascending order.
func sortLst(arr []artist, sortCriteria string) {
	fmt.Println("effectively used")
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
