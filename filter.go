package main

import (
	"strconv"
	"strings"
)

func checkMembersNb(artist artistData, membersNb string) bool {
	nb, _ := strconv.Atoi(membersNb)
	return len(artist.Members) == nb
}

func compareFADate(artist artistData, date string) bool {
	if strings.Compare(artist.FirstAlbum, date) == 0 {
		return true
	}
	return false
}

func filterArtists(arr []artistData, membersNbs []string) []artistData {
	newArr := []artistData{}
	for _, membersNb := range membersNbs {
		tempFiltered := filter(arr, membersNb, checkMembersNb)
		newArr = append(newArr, tempFiltered...)
	}
	return newArr
}

func filter(arr []artistData, criteria string, check func (artistData, string) bool) []artistData {
	newArr := []artistData{}
	for _, artist := range arr {
		if check(artist, criteria) {
			newArr = append(newArr, artist)
		}
	}
	return newArr
}