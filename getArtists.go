package main

import (
	"strings"
)

// getArtistsData() gets the API data and makes them presentable
func getArtistsData() *errorPage {
	err := getAPIData()
	if err != nil {
		return err
	}
	for i := 0; i < len(artistsLst); i++ {
		artistsLst[i].MembersCount = len(artistsLst[i].Members)
		artistsLst[i].LocDate = getRelPretty(rels.Lst[i].DatesLocations)
		artistsLst[i].LocCount = getLocCount(locs.Lst[i].Locations)
		artistsLst[i].Performances = len(dates.Lst[i].Dates)
	}
	return err
}

// getLocCount() uses locLst to count
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

// getRelPretty() makes the key value in the original Relations List presentable
func getRelPretty(ogRelLst map[string][]string) map[string][]string {
	newRelLst := make(map[string][]string)

	for key, value := range ogRelLst {

		newKey := strings.ReplaceAll(key, "_", " ")
		newKey = strings.ReplaceAll(newKey, "-", ", ")
		newKey = strings.ToUpper(newKey)

		newRelLst[newKey] = value
	}
	return newRelLst
}
