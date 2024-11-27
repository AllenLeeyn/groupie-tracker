package main

type artistData struct {
	Image        string
	Name         string
	Members      []string
	MembersCount int
	CreationDate int
	FirstAlbum   string
	Locations    []string
	Performances int
	Relations    map[string][]string
}

var artistsData []artistData

func getArtist() {
	grabAPI()
	for i := range artists {
		artistsData = append(artistsData, artistData{
			Image:        artists[i].Image,
			Name:         artists[i].Name,
			Members:      artists[i].Members,
			MembersCount: len(artists[i].Members),
			CreationDate: artists[i].CreationDate,
			FirstAlbum:   artists[i].FirstAlbum,
			Locations:    locLst.Locations[i].Locations,
			Performances: len(dateLst.Dates[i].Dates),
			Relations:    relLst.Relations[i].DatesLocations,
		})
	}
}
