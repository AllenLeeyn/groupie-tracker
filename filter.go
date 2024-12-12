package main

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"
	"strings"
)

func getFiltrs(req *http.Request) filters {
	filtrs := filters{}
	filtrs.NbChecked = req.Form["members number"]

	filtrs.ApplyCreationDFltr = req.Form.Get("applyRange")
	filtrs.CreateDate = checkDate(req.Form.Get("range"))
	filtrs.CreateRange, _ = strconv.Atoi(req.Form.Get("years range"))

	filtrs.Locations = checkLocations(req.Form.Get("locations"))

	filtrs.ApplyFirstADFltr = req.Form.Get("applyRange0")
	filtrs.FirstDate = checkDate(req.Form.Get("range0"))
	filtrs.FirstRange, _ = strconv.Atoi(req.Form.Get("yearsrange0"))

	return filtrs
}

func checkDate(date string) int {
	if len(date) == 4 {
		result, _ := strconv.Atoi(date)
		return result
	}
	return 0
}

func checkLocations(locationsArr string) []string {
	if locationsArr == "" {
		return []string{}
	}
	return strings.Split(locationsArr, "\n")
}

func mainFilter(req *http.Request) error {
	var newArtistsLst []artist = slices.Clone(artistsLst)

	if err := req.ParseForm(); err != nil {
		return BadRequestErr
	}
	filtrs := getFiltrs(req)

	if len(req.Form["submit button"]) == 1 && req.Method != "POST" {
		return MethodNotAllowedErr
	}
	if len(filtrs.NbChecked) != 0 {
		newArtistsLst = filterMembersNbs(newArtistsLst, filtrs.NbChecked)
	}
	if len(filtrs.Locations) != 0 {
		newArtistsLst = filterLocations(newArtistsLst, filtrs.Locations)
	}
	if filtrs.ApplyFirstADFltr == "on" {
		newArtistsLst = filterDates(newArtistsLst, filtrs.FirstDate,
			filtrs.FirstRange, 1963, 2018, compareFADate)
	}
	if filtrs.ApplyCreationDFltr == "on" {
		newArtistsLst = filterDates(newArtistsLst, filtrs.CreateDate,
			filtrs.CreateRange, 1958, 2015, compareCreationDate)
	}

	homePage.Artists = newArtistsLst
	homePage.Filters = filtrs
	return nil
}

func filterMembersNbs(arr []artist, membersNbs []string) []artist {
	newArr := []artist{}
	for _, membersNb := range membersNbs {
		tempFiltered := filter(arr, membersNb, compareMembersNb)
		newArr = append(newArr, tempFiltered...)
	}
	return newArr
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

func filterDates(arr []artist, date, rng, min, max int,
	check func(artist, string) bool) []artist {
	newArr := []artist{}
	startDate, endDate := getYearsRange(date, rng, min, max)
	for start := startDate; start <= endDate; start++ {
		newArr = append(newArr, filter(arr, strconv.Itoa(start), check)...)
	}
	return newArr
}

func getYearsRange(date, rng, min, max int) (startDate, endDate int) {
	if date < min || date > max {
		date = min
	}
	startDate, endDate = date, date+rng
	if rng < 0 {
		startDate, endDate = endDate, startDate
	}
	if startDate < min {
		startDate = min
	}
	if endDate > max {
		endDate = max
	}
	return
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

func compareMembersNb(artist artist, membersNb string) bool {
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

func compareCreationDate(artist artist, creationDate string) bool {
	intDate, _ := strconv.Atoi(creationDate)
	return artist.CreationDate == intDate
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

func isArtistInArr(arr []artist, name string) bool {
	for _, artist := range arr {
		if artist.Name == name {
			return true
		}
	}
	return false
}
