package main

import (
	"strings"
)

type artistData struct {
	Index          int
	Image          string
	Name           string
	Members        []string
	MembersCount   int
	CreationDate   int
	FirstAlbum     string
	Locations      map[string][]string
	LocationsCount int
	Performances   int
	Relations      map[string][]string
}

var artistsData []artistData

func getArtist() {
	grabAPI()
	for i := range artists {
		artistsData = append(artistsData, artistData{
			Index:          artists[i].Id,
			Image:          artists[i].Image,
			Name:           artists[i].Name,
			Members:        artists[i].Members,
			MembersCount:   len(artists[i].Members),
			CreationDate:   artists[i].CreationDate,
			FirstAlbum:     artists[i].FirstAlbum,
			Locations:      getRelationsClean(relLst.Relations[i].DatesLocations),
			LocationsCount: getLocCount(locLst.Locations[i].Locations),
			Performances:   len(dateLst.Dates[i].Dates),
			Relations:      relLst.Relations[i].DatesLocations,
		})
	}
}

func getLocCount(locations []string) int {

	locLst := make(map[string]string)
	for _, loc := range locations {
		splited := strings.Split(loc, "-")

		if len(splited) == 2 {
			locLst[splited[1]] = splited[0]
		}
	}
	return len(locLst)
}

func getRelationsClean(ogRelLst map[string][]string) map[string][]string {
	newRelLst := make(map[string][]string)

	for key, value := range ogRelLst {
		newKey := strings.ReplaceAll(strings.ReplaceAll(key, "_", " "), "-", ", ")
		newKey = capitalize(newKey)

		newRelLst[newKey] = value
	}

	return newRelLst
}

func capitalize(s string) (result string) {
	splited := strings.Split(s, " ")
	for _, word := range splited {
		if word == "usa" || word == "uk" {
			word = strings.ToUpper(word)
		}
		result += strings.ToUpper(string(word[0])) + word[1:] + " "
	}
	return result[:len(result)-1]
}
