package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"text/template"
)

const apiURL = "https://groupietrackers.herokuapp.com/api/"

type listPage struct {
	Artists []artist
	SortBy  string
	Order   string

	// filters
	NbChecked  []string
	DateFA     string
	Locations  []string
	EarliestDt string
	LatestDt   string
	ApplyRange string
}

var homePage *listPage = &listPage{}

type artistPage struct {
	Artist artist
}

type errorPage struct {
	errorCode int
	errorMsg  string
}

func (e errorPage) Error() string {
	return e.errorMsg
}

var indexTmpl = template.Must(template.ParseFiles("templates/index.html"))
var artistTmpl = template.Must(template.ParseFiles("templates/artist.html"))

func main() {
	getArtistsData()
	http.Handle("/static/", http.FileServer(http.Dir("assets/")))
	http.HandleFunc("/", homeHandler)

	port := "localhost:8081"
	log.Println("Listening on http://" + port)
	http.ListenAndServe(port, nil)
}

func checkFADate(date []string) error {
	if len(date) == 0 {
		return errors.New("invalid date format")
	}
	isMatch, _ := regexp.Match("([0-9]{2}-[0-9]{2}-[0-9]{4})", []byte(date[0]))
	if len(date[0]) != 10 || !isMatch {
		return errors.New("invalid date format")
	}
	return nil
}

func checkGetFADate(dateArr []string) string {
	var date string

	err := checkFADate(dateArr)
	if err == nil {
		date = dateArr[0]
	} else {
		date = ""
	}
	return date
}

func checkGetLocations(locationsArr []string) []string {
	if len(locationsArr) == 0 {
		return []string{}
	}
	return strings.Fields(locationsArr[0])
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

// arrangeArtists checks for method error and sorts/filter artists
// filter by membersNb, dataFA
func arrangeArtists(w http.ResponseWriter, req *http.Request) error {
	// var homePage *listPage
	var newArtistsLst []artist = slices.Clone(artistsLst)

	checkErr(req.ParseForm())
	// make a new copy, so when the artists' array is filtered, the artists
	// who are not displayed aren't lost for the whole execution of the program
	membersNb := req.Form["members number"]
	dateFA := checkGetFADate(req.Form["first album date"]) // FA: first album
	locations := checkGetLocations(req.Form["locations"])
	rangeValue := req.Form["range"]
	applyRange := req.Form["applyRange"]
	fmt.Println("range value:", rangeValue)
	if (len(req.Form["submit button"]) == 1 && req.Method != "POST") ||
		((len(req.Form["submit button"]) == 0 && len(req.Form["sort"]) == 0 && len(req.Form["switch-order"]) == 0) && req.Method != "GET") {
		return errorPage{405, "405 method not allowed"}
	}
	if len(membersNb) != 0 {
		newArtistsLst = filterArtists(newArtistsLst, membersNb)
	}
	if dateFA != "" {
		newArtistsLst = filter(newArtistsLst, dateFA, compareFADate)
	}
	if len(locations) != 0 {
		newArtistsLst = filterLocations(newArtistsLst, locations)
	}
	if len(applyRange) != 0 {
		newArtistsLst = filter(newArtistsLst, rangeValue[0], compareCreationDate)
	}
	// sortAlph(newArtistsLst)
	order, sortCriteria, err := sortArtists(w, req, newArtistsLst)
	checkErr(err)
	homePage = &listPage{
		Artists:    newArtistsLst,
		NbChecked:  membersNb,
		DateFA:     dateFA,
		Locations:  locations,
		Order:      order,
		SortBy:     sortCriteria,
		EarliestDt: strconv.Itoa(earliestCreationDate(newArtistsLst)),
		LatestDt:   strconv.Itoa(latestCreationDate(newArtistsLst)),
	}
	// homePage.Artists = newArtistsLst
	// homePage.NbChecked = membersNb
	// homePage.DateFA = dateFA
	// homePage.Locations = locations
	// homePage.Order = order
	// homePage.SortBy = sortCriteria
	return nil
}

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
	// if req.Method != "POST" {
	// 	newArtistsLst = slices.Clone(artistsLst)
	// }

	err := arrangeArtists(w, req)
	if err != nil {
		errPage := err.(errorPage) // type assertion
		http.Error(w, errPage.errorMsg, errPage.errorCode)
		return
	}
	indexTmpl.Execute(w, homePage)
}

func earliestCreationDate(arr []artist) int {
	if len(arr) == 0 {
		return -1
	}
	earliestDate := arr[0].CreationDate
	for _, artist := range arr {
		if earliestDate > artist.CreationDate {
			earliestDate = artist.CreationDate
		}
	}
	return earliestDate
}

func latestCreationDate(arr []artist) int {
	if len(arr) == 0 {
		return -1
	}
	var latestDate int
	for _, artist := range arr {
		if latestDate < artist.CreationDate {
			latestDate = artist.CreationDate
		}
	}
	return latestDate
}
