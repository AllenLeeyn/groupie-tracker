package main

import (
	"net/http"
)

type filters struct {
	NbChecked  []string
	DateFA     string
	Locations  []string
	EarliestDt string
	LatestDt   string
	ApplyRange string
}

type listPage struct {
	Artists []artist
	SortBy  string
	Order   string
	Filters filters
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
	if ArtistErr != nil {
		errorHandler(&w, InternalServerErr)
		return
	}

	for index, artistEntry := range artistsLst {
		if req.URL.Path[1:] == artistEntry.Name &&
			req.Method == http.MethodGet {
			artistTmpl.Execute(w, struct{ Artist artist }{Artist: artistsLst[index]})
			return
		}
	}
	if req.URL.Path != "/" {
		errorHandler(&w, NotFoundErr)
		return
	}

	if err := mainFilter(req); err != nil {
		errPage := err.(errorPage) // type assertion
		errorHandler(&w, errPage)
		return
	}
	if err := getSortedArtists(req); err != nil {
		errPage := err.(errorPage) // type assertion
		errorHandler(&w, errPage)
		return
	}
	indexTmpl.Execute(w, homePage)
}

func errorHandler(w *http.ResponseWriter, err errorPage) {
	(*w).WriteHeader(err.ErrorCode)
	errTmpl.Execute(*w, err)
}
