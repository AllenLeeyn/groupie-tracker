package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"slices"
	"strings"
	"text/template"
)

const apiURL = "https://groupietrackers.herokuapp.com/api/"

type listPage struct {
	Artists []artistData

	// filters
	NbChecked []string
	DateFA    string
	Locations []string
}

type artistPage struct {
	Artist artistData
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
	getArtist()
	http.Handle("/static/", http.FileServer(http.Dir("assets/")))

	http.HandleFunc("/", homeHandler)

	port := "localhost:8081"
	log.Println("Listening on " + port)
	http.ListenAndServe(port, nil)
}

func checkFADate(date []string) error {
	fmt.Println(date)
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
	fmt.Println(date)
	return date
}

func checkGetLocations(locationsArr []string) []string {
	if len(locationsArr) == 0 {
		return []string{}
	}
	return strings.Fields(locationsArr[0])
}

// arrangeArtists checks for method error and sorts/filter artists
func arrangeArtists(req *http.Request) (*listPage, error) {
	var homePage *listPage

	newArtistsDt := slices.Clone(artistsData)
	checkErr(req.ParseForm())
	membersNb := req.Form["members number"]
	dateFA := checkGetFADate(req.Form["first album date"]) // FA: first album
	locations := checkGetLocations(req.Form["locations"])
	if (len(req.Form["submit button"]) == 1 && req.Method != "POST") ||
		(len(req.Form["submit button"]) == 0 && req.Method != "GET") {
		return nil, errorPage{405, "405 method not allowed"}
	}
	if len(membersNb) != 0 {
		newArtistsDt = filterArtists(artistsData, membersNb)
	}
	if dateFA != "" {
		newArtistsDt = filter(newArtistsDt, dateFA, compareFADate)
	}
	if len(locations) != 0 {
		fmt.Println("locations:", locations)
		newArtistsDt = filterLocations(newArtistsDt, locations)
	}
	sortAlph(newArtistsDt)
	homePage = &listPage{
		Artists:   newArtistsDt,
		NbChecked: membersNb,
		DateFA:    dateFA,
		Locations: locations,
	}
	fmt.Println("newArtistsDt:", newArtistsDt)
	fmt.Println("artistsDt:", artistsData)
	return homePage, nil
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	// fmt.Println("ARTISTSDATA:", artistsData)
	for index, artist := range artists {
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

	homePage, err := arrangeArtists(req)
	if err != nil {
		errPage := err.(errorPage) // type assertion
		http.Error(w, errPage.errorMsg, errPage.errorCode)
		return
	}
	indexTmpl.Execute(w, homePage)
}

func artistHandler(w http.ResponseWriter, index int) {
	artPage := &artistPage{Artist: artistsData[index]}

	artistTmpl.Execute(w, artPage)
}
