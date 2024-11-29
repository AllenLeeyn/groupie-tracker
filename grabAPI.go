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
func getAPIData() {
	dataNames := []string{"artists", "locations", "dates", "relation"}

	for _, dataName := range dataNames {
		apiResp, err := http.Get(apiURL + dataName)
		checkErr(err)
		defer apiResp.Body.Close()

		if apiResp.StatusCode != http.StatusOK {
			log.Fatalf("Error: %v", apiResp.StatusCode)
		}
		apiRaw, err := io.ReadAll(apiResp.Body)
		checkErr(err)

		switch dataName {
		case "artists":
			err := json.Unmarshal(apiRaw, &artistsLst)
			checkErr(err)
		case "locations":
			err := json.Unmarshal(apiRaw, &locs)
			checkErr(err)
		case "dates":
			err := json.Unmarshal(apiRaw, &dates)
			checkErr(err)
		case "relation":
			err := json.Unmarshal(apiRaw, &rels)
			checkErr(err)
		}
	}
	checkAPIData()
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
		locCount := len(rels.Lst[i].DatesLocations)
		datesCount := func() (count int) {
			for _, dates := range rels.Lst[i].DatesLocations {
				for range dates {
					count++
				}
			}
			return
		}()
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
