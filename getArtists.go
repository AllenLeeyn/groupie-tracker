package main

import (
	"strings"
)

// getArtistsData() gets the API data and makes them presentable
func getArtistsData() {
	getAPIData()
	for i := 0; i < len(artistsLst); i++ {
		artistsLst[i].MembersCount = len(artistsLst[i].Members)
		artistsLst[i].LocDate = getRelClean(rels.Lst[i].DatesLocations)
		artistsLst[i].LocCount = getLocCount(locs.Lst[i].Locations)
		artistsLst[i].Performances = len(dates.Lst[i].Dates)
	}
}

// getLocCount uses locLst to count
// how many countries an artist has performed in
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

// getRelClean() makes the key value in the original Relations List presentable
func getRelClean(ogRelLst map[string][]string) map[string][]string {
	newRelLst := make(map[string][]string)

	capitalize := func(words []string) (result string) {
		for i := 0; i < len(words); i++ {
			if words[i] == "usa" || words[i] == "uk" {
				words[i] = strings.ToUpper(words[i])
			} else {
				words[i] = strings.ToTitle(words[i])
			}
		}
		return strings.Join(words, " ")
	}
	for key, value := range ogRelLst {
		newKey := strings.ReplaceAll(key, "_", " ")
		newKey = strings.ReplaceAll(newKey, "-", ", ")
		newKey = capitalize(strings.Split(newKey, " "))
		newRelLst[newKey] = value
	}
	return newRelLst
}
