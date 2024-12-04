package main

import (
	"encoding/json"
	"io"
	"net/http"
)

const apiURL = "https://groupietrackers.herokuapp.com/api/"

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

		apiRaw, err := io.ReadAll(apiResp.Body)
		defer apiResp.Body.Close()
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
	return nil
}
