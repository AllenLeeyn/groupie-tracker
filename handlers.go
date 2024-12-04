package main

import (
	"fmt"
	"net/http"
	"slices"
	"sort"
)

var (
	NotFoundErr = errorPage{
		ErrorCode: 404,
		ErrorMsg:  "Oops! The page you're looking for doesn't exist.",
	}
	InternalServerErr = errorPage{
		ErrorCode: 500,
		ErrorMsg:  "Internal Server Error",
	}
	BadRequestErr = errorPage{
		ErrorCode: 400,
		ErrorMsg:  "Bad Request",
	}
	MethodNotAllowedErr = errorPage{
		ErrorCode: http.StatusMethodNotAllowed,
		ErrorMsg:  "Method Not Allowed",
	}
	BadGatewayErr = errorPage{
		ErrorCode: http.StatusBadGateway,
		ErrorMsg:  "Bad Gateway",
	}
)

// arrangeArtists checks for method error and sorts/filter artists
func arrangeArtists(req *http.Request) error {
	// var homePage *listPage
	var newArtistsLst []artist = slices.Clone(artistsLst)

	checkErr(req.ParseForm())
	order, sortCriteria, err := sortArtists(req, newArtistsLst)
	// // checkErr(err)
	if err != nil {
		return err
	}
	fmt.Println(err)
	homePage = &listPage{
		Artists: newArtistsLst,
		Order:   order,
		SortBy:  sortCriteria,
	}
	return nil
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	if err2 != nil {
		http.Error(w, "500 Internal Server Error", 500)
		return
	}
	if err1 != nil {
		w.WriteHeader(InternalServerErr.ErrorCode)
		errTmpl.Execute(w, InternalServerErr)
		return
	} else if ArtistErr != nil {
		w.WriteHeader(InternalServerErr.ErrorCode)
		errTmpl.Execute(w, *ArtistErr)
		return
	}

	for index, artist := range artistsLst {
		if req.URL.Path[1:] == artist.Name &&
			req.Method == http.MethodGet {
			artistHandler(w, index)
			return
		}
	}
	if req.URL.Path != "/" {
		w.WriteHeader(404)
		errTmpl.Execute(w, NotFoundErr)
		return
	}

	err := arrangeArtists(req)
	if err != nil {
		errPage := err.(errorPage) // type assertion
		w.WriteHeader(errPage.ErrorCode)
		errTmpl.Execute(w, errPage)
		return
	}
	indexTmpl.Execute(w, homePage)
}

// artistHandler() generates the html response for an artist
func artistHandler(w http.ResponseWriter, index int) {
	if artistTmplErr != nil {
		w.WriteHeader(InternalServerErr.ErrorCode)
		errTmpl.Execute(w, InternalServerErr)
		return
	}
	artPage := &artistPage{Artist: artistsLst[index]}
	artistTmpl.Execute(w, artPage)
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
