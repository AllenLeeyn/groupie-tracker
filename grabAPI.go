package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	MembersCount int
	LocDate      map[string][]string
	LocCount     int
	Performances int
}

var artistsLst []artist

var locs struct {
	Lst []struct {
		Id        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

var dates struct {
	Lst []struct {
		Id    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

var rels struct {
	Lst []struct {
		Id             int                 `json:"id"`
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

// getAPIData() sends a http GET request to API and
// unmarshal the json data and store them in their corresponding struct
func getAPIData() *errorPage {
	dataNames := []string{"artists", "locations", "dates", "relation"}

	for _, dataName := range dataNames {
		var err error
		apiResp, err := http.Get(apiURL + dataName)
		if err != nil {
			return &BadGatewayErr
		}
		defer apiResp.Body.Close()

		apiRaw, err := io.ReadAll(apiResp.Body)
		if err != nil {
			return &BadGatewayErr
		}

		switch dataName {
		case "artists":
			err = json.Unmarshal(apiRaw, &artistsLst)
		case "locations":
			err = json.Unmarshal(apiRaw, &locs)
		case "dates":
			err = json.Unmarshal(apiRaw, &dates)
		case "relation":
			err = json.Unmarshal(apiRaw, &rels)
		}
		if err != nil {
			return &NotFoundErr
		}
	}
	checkAPIData()
	return nil
}

// checkAPIData() does simple check to see if the data matches
func checkAPIData() {
	artistsCount := len(artistsLst)
	if len(locs.Lst) != artistsCount ||
		len(dates.Lst) != artistsCount ||
		len(rels.Lst) != artistsCount {
		log.Fatal("ERROR: Entry count does not tally")
	}

	for i := range artistsCount {
		locCount, datesCount := len(rels.Lst[i].DatesLocations), 0
		for _, dates := range rels.Lst[i].DatesLocations {
			datesCount += len(dates)
		}

		if locCount != len(locs.Lst[i].Locations) ||
			datesCount != len(dates.Lst[i].Dates) {
			log.Printf("ERROR: Entry [%v] does not tally\n", i)
			log.Printf("relations %v v locations %v\n", locCount, len(locs.Lst[i].Locations))
			log.Printf("relations %v v dates %v\n", datesCount, len(dates.Lst[i].Dates))
			log.Println("=========================================")
		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}
}
