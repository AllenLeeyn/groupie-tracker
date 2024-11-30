package main

import (
	"strconv"
)

func checkMembersNb(artist artistData, membersNb string) bool {
	nb, _ := strconv.Atoi(membersNb)
	return len(artist.Members) == nb
}

func compareLoc(artist artistData, loc string) bool {
	for _, artistLoc := range artist.Locations {
		if artistLoc == loc {
			return true
		}
	}
	return false
}

func compareFADate(artist artistData, date string) bool {
	return artist.FirstAlbum == date
}

func filterArtists(arr []artistData, membersNbs []string) []artistData {
	newArr := []artistData{}
	for _, membersNb := range membersNbs {
		tempFiltered := filter(arr, membersNb, checkMembersNb)
		newArr = append(newArr, tempFiltered...)
	}
	return newArr
}

// differenceElements returns array of artists that are in a which are not in b
func differenceElements(a []artistData, b []artistData) []artistData {
	difference := []artistData{}

	for _, artist := range a {
		if !isArtistInArr(b, artist.Name) {
			difference = append(difference, artist)
		}
	}
	return difference
}

func filterLocations(arr []artistData, locations []string) []artistData {
	newArr := []artistData{}
	for _, loc := range locations {
		tempFiltered := filter(arr, loc, compareLoc)
		tempFiltered = differenceElements(tempFiltered, newArr)
		newArr = append(newArr, tempFiltered...)
	}
	return newArr
}

func filter(arr []artistData, criteria string, check func(artistData, string) bool) []artistData {
	newArr := []artistData{}
	for _, artist := range arr {
		if check(artist, criteria) && !isArtistInArr(newArr, artist.Name) {
			newArr = append(newArr, artist)
		}
	}
	return newArr
}

func isArtistInArr(arr []artistData, name string) bool {
	for _, artist := range arr {
		if artist.Name == name {
			return true
		}
	}
	return false
}
