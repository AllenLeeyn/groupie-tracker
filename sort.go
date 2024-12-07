package main

import (
	"net/http"
	"sort"
)

// getSortedArtists() checks for method error and sorts artists
func getSortedArtists(req *http.Request) error {
	// var newArtistsLst []artist = slices.Clone(artistsLst)

	order, sortCriteria, err := sortArtists(req, homePage.Artists)
	// homePage = &listPage{
	// 	// Artists: newArtistsLst,
	// 	Order:   order,
	// 	SortBy:  sortCriteria,
	// }
	homePage.Order = order
	homePage.SortBy = sortCriteria
	return err
}

var order = "▼"

func sortArtists(req *http.Request, arr []artist) (string, string, error) {
	var sortCriteria string

	// POST method is used for sorting list.
	// Invalid request is ignored and default setting is used.
	if req.Method == http.MethodPost {
		sortCriteria = req.FormValue("sort")
		if sortCriteria != "creation_date" && sortCriteria != "name" {
			sortCriteria = "default"
		}
		sortLst(arr, sortCriteria)

		pageOrder := req.FormValue("switch-order")
		if pageOrder == "▲" {
			order = "▼"
		} else if pageOrder == "▼" || order == "▲" {
			order = "▲"
			revLst(arr)
		}
	} else if req.Method != http.MethodGet {
		return "", "", MethodNotAllowedErr
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
