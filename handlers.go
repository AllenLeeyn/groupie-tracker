package main

import (
	"net/http"
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

// homeHandler() handles all request and url.
// passes '/artistName' to artistHandler().
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

	err := getSortedArtists(req)

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
