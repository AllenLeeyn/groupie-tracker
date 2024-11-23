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
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"dates"`
	Relations    string   `json:"relations"`
}

var artists []artist

type locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

var locLst struct {
	Locations []locations `json:"index"`
}

type dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

var dateLst struct {
	Dates []dates `json:"index"`
}

type relations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

var relLst struct {
	Relations []relations `json:"index"`
}

func grabAPI() {
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
			err := json.Unmarshal(apiRaw, &artists)
			checkErr(err)
		case "locations":
			err := json.Unmarshal(apiRaw, &locLst)
			checkErr(err)
		case "dates":
			err := json.Unmarshal(apiRaw, &dateLst)
			checkErr(err)
		case "relation":
			err := json.Unmarshal(apiRaw, &relLst)
			checkErr(err)
		}
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}
}
