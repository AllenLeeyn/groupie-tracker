package main

import (
	"net/http"
)

type listPage struct {
	Artists []artist
	SortBy  string
	Order   string
}

var homePage *listPage = &listPage{}

type errorPage struct {
	ErrorCode int
	ErrorMsg  string
}

func (e errorPage) Error() string {
	return e.ErrorMsg
}

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
		errorHandler(&w, InternalServerErr)
		return
	} else if ArtistErr != nil {
		errorHandler(&w, InternalServerErr)
		return
	}

	for index, artist := range artistsLst {
		if req.URL.Path[1:] == artist.Name &&
			req.Method == http.MethodGet {
			artistHandler(&w, index)
			return
		}
	}
	if req.URL.Path != "/" {
		errorHandler(&w, NotFoundErr)
		return
	}

	if err := getSortedArtists(req); err != nil {
		errPage := err.(errorPage) // type assertion
		errorHandler(&w, errPage)
		return
	}
	indexTmpl.Execute(w, homePage)
}

// artistHandler() generates the html response for an artist
func artistHandler(w *http.ResponseWriter, index int) {
	if artistTmplErr != nil {
		errorHandler(w, InternalServerErr)
		return
	}
	artistTmpl.Execute(*w, struct{ Artist artist }{Artist: artistsLst[index]})
}

func errorHandler(w *http.ResponseWriter, err errorPage) {
	(*w).WriteHeader(err.ErrorCode)
	errTmpl.Execute(*w, err)
}
