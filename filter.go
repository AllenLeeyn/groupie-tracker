package main

import (
	"errors"
	"net/http"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func getFiltrs(req *http.Request) (filters, string, []string) {
	filtrs := filters{}
	filtrs.NbChecked = req.Form["members number"]
	filtrs.DateFA = checkGetFADate(req.Form["first album date"]) // FA: first album
	filtrs.Locations = checkGetLocations(req.Form["locations"])
	CreationD := req.Form.Get("range")
	x, _ := strconv.Atoi(CreationD)
	if x == 0 {
		CreationD = ""
	}
	applyCreationDFltr := req.Form["applyRange"]
	firstADate := req.Form["range0"]
	applyFirstADFltr := req.Form["applyRange0"]

	if len(applyFirstADFltr) != 0 {
		if len(req.Form["yearsrange0"]) != 0 {
			filtrs.YearsRange0 = req.Form["yearsrange0"][0]
		}
		filtrs.ApplyFirstADFltr = applyFirstADFltr[0]
	}
	if len(applyCreationDFltr) != 0 {
		if len(req.Form["years range"]) != 0 {
			filtrs.YearsRange = req.Form["years range"][0]
		}
		filtrs.ApplyCreationDFltr = applyCreationDFltr[0]
	}

	return filtrs, CreationD, firstADate
}

func mainFilter(req *http.Request) error {
	var newArtistsLst []artist = slices.Clone(artistsLst)

	if err := req.ParseForm(); err != nil {
		return BadRequestErr
	}

	filtrs, CreationD, firstADate := getFiltrs(req)

	if len(req.Form["submit button"]) == 1 && req.Method != "POST" {
		return MethodNotAllowedErr
	}
	if len(filtrs.NbChecked) != 0 {
		newArtistsLst = filterArtists(newArtistsLst, filtrs.NbChecked)
	}
	if len(filtrs.Locations) != 0 {
		newArtistsLst = filterLocations(newArtistsLst, filtrs.Locations)
	}
	if filtrs.ApplyFirstADFltr == "on" {
		// newArtistsLst = filterFirstADate(newArtistsLst, firstADate[0], filtrs.YearsRange0)
		newArtistsLst = filterRange(newArtistsLst, firstADate[0], filtrs.YearsRange0, 1963, 2018, compareFADate)
	}
	if filtrs.ApplyCreationDFltr == "on" {
		// newArtistsLst = filterCreationDateRange(newArtistsLst, CreationD[0], filtrs.YearsRange)
		// newArtistsLst = filterCreationDateRange(newArtistsLst, , filtrs.YearsRange)
		newArtistsLst = filterRange(newArtistsLst, CreationD, filtrs.YearsRange, 1958, 2015, compareCreationDate)
	}

	homePage.Artists = newArtistsLst
	homePage.Filters = filtrs
	return nil
}

func getYearsRange(rangeValue, yearsStr string, minYears, maxYears int) (startDate string, endDate string) {
	yearsNb, _ := strconv.Atoi(yearsStr)
	rangeValueNb, _ := strconv.Atoi(rangeValue)
	if yearsNb < 0 {
		start := rangeValueNb + yearsNb
		if start < minYears {
			start = minYears
		}
		startDate = strconv.Itoa(start)
		endDate = rangeValue
	} else {
		end := rangeValueNb + yearsNb
		if end > maxYears {
			end = maxYears
		}
		startDate = rangeValue
		endDate = strconv.Itoa(end)
	}
	return
}

func filterRange(arr []artist, rangeValue, yearsRange string, minYear, maxYear int,
	compare func(artist, string) bool) []artist {
	newArr := []artist{}
	startDate, endDate := getYearsRange(rangeValue, yearsRange, minYear, maxYear)
	end, _ := strconv.Atoi(endDate)
	for start, _ := strconv.Atoi(startDate); start <= end; start++ {
		newArr = append(newArr, filter(arr, strconv.Itoa(start), compare)...)
	}
	return newArr
}

func checkFADate(date []string) error {
	if len(date) == 0 {
		return errors.New("invalid date format")
	}
	isMatch, _ := regexp.Match("([0-9]{4})", []byte(date[0]))
	if !isMatch {
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
	if len(locationsArr) == 0 || (len(locationsArr) == 1 && locationsArr[0] == "") {
		return []string{}
	}
	return strings.Split(locationsArr[0], "\n")
}

func checkMembersNb(artist artist, membersNb string) bool {
	nb, _ := strconv.Atoi(membersNb)
	return len(artist.Members) == nb
}

func compareLoc(artist artist, loc string) bool {
	for artistLoc := range artist.LocDate {
		if artistLoc == loc {
			return true
		}
	}
	return false
}

func compareFADate(artist artist, date string) bool {
	if len(date) < len(artist.FirstAlbum) {
		return artist.FirstAlbum[len(artist.FirstAlbum)-len(date):] == date
	}
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
