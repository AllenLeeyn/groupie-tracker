package main

import (
	"fmt"
	"strconv"
	"strings"
)

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
	// fmt.Println("LocDate:", arr[0].)
	fmt.Println("---------------------")
	// spew.Dump(arr[0].LocDate)
	fmt.Println("---------------------")
	newArr := []artist{}
	for _, loc := range locations {
		tempFiltered := filter(arr, loc, compareLoc)
		tempFiltered = differenceElements(tempFiltered, newArr)
		newArr = append(newArr, tempFiltered...)
	}
	return newArr
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
