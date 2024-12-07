package main

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func mainFilter(req *http.Request) error {
	var newArtistsLst []artist = slices.Clone(artistsLst)

	if err := req.ParseForm(); err != nil {
		return BadRequestErr
	}

	filtrs := filters{}
	filtrs.NbChecked = req.Form["members number"]
	filtrs.DateFA = checkGetFADate(req.Form["first album date"]) // FA: first album
	filtrs.Locations = checkGetLocations(req.Form["locations"])
	rangeValue := req.Form["range"]
	applyRange := req.Form["applyRange"]
	fmt.Println("range value:", rangeValue)
	if (len(req.Form["submit button"]) == 1 && req.Method != "POST") ||
		((len(req.Form["submit button"]) == 0 && len(req.Form["sort"]) == 0 && len(req.Form["switch-order"]) == 0) && req.Method != "GET") {
		// return errorPage{405, "405 method not allowed"}
		return MethodNotAllowedErr
	}
	if len(filtrs.NbChecked) != 0 {
		newArtistsLst = filterArtists(newArtistsLst, filtrs.NbChecked)
	}
	if filtrs.DateFA != "" {
		newArtistsLst = filter(newArtistsLst, filtrs.DateFA, compareFADate)
	}
	if len(filtrs.Locations) != 0 {
		newArtistsLst = filterLocations(newArtistsLst, filtrs.Locations)
	}
	if len(applyRange) != 0 {
		newArtistsLst = filter(newArtistsLst, rangeValue[0], compareCreationDate)
		filtrs.ApplyRange = applyRange[0]
	}
	filtrs.EarliestDt = strconv.Itoa(earliestCreationDate(newArtistsLst))
	filtrs.LatestDt = strconv.Itoa(latestCreationDate(newArtistsLst))

	homePage.Artists = newArtistsLst
	homePage.Filters = filtrs
	return nil
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

func checkMembersNb(artist artist, membersNb string) bool {
	nb, _ := strconv.Atoi(membersNb)
	return len(artist.Members) == nb
}

func compareLoc(artist artist, loc string) bool {
	for artistLoc := range artist.LocDate {
		if strings.Contains(artistLoc, loc) {
			return true
		}
	}
	return false
}

func compareFADate(artist artist, date string) bool {
	return artist.FirstAlbum == date
}

func filterArtists(arr []artist, membersNbs []string) []artist {
	newArr := []artist{}
	for _, membersNb := range membersNbs {
		tempFiltered := filter(arr, membersNb, checkMembersNb)
		newArr = append(newArr, tempFiltered...)
	}
	return newArr
}

// differenceElements returns array of artists that are in a which are not in b
func differenceElements(a []artist, b []artist) []artist {
	difference := []artist{}

	for _, artist := range a {
		if !isArtistInArr(b, artist.Name) {
			difference = append(difference, artist)
		}
	}
	return difference
}

func filterLocations(arr []artist, locations []string) []artist {
	newArr := []artist{}
	for _, loc := range locations {
		tempFiltered := filter(arr, loc, compareLoc)
		tempFiltered = differenceElements(tempFiltered, newArr)
		newArr = append(newArr, tempFiltered...)
	}
	return newArr
}

func compareCreationDate(artist artist, creationDate string) bool {
	intDate, _ := strconv.Atoi(creationDate)
	return artist.CreationDate == intDate
}

func filter(arr []artist, criteria string, check func(artist, string) bool) []artist {
	newArr := []artist{}
	for _, artist := range arr {
		if isArtistInArr(newArr, artist.Name) {
			fmt.Println("?0")
		}
		if check(artist, criteria) && !isArtistInArr(newArr, artist.Name) {
			newArr = append(newArr, artist)
		}
	}
	return newArr
}

func isArtistInArr(arr []artist, name string) bool {
	for _, artist := range arr {
		if artist.Name == name {
			return true
		}
	}
	return false
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
