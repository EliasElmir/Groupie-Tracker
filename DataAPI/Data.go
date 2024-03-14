package DataAPI

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const URL = "https://groupietrackers.herokuapp.com/api"

type Artist []DataArtist
type LocationIndex struct {
	Location []DataLocations `json:"index"`
}

type DatesIndex struct {
	Dates []DataDate `json:"index"`
}

type RelationIndex struct {
	Relation []DataRelations `json:"index"`
}

type DataArtist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    []string
	ConcertDates []string
	Relations    map[string][]string
}

type DataLocations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     []string `json:"dates"`
}

type DataDate struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DataRelations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetArtistData() (Artist, error) {
	var Data Artist
	resArtist, err := http.Get(URL + "/artists")
	if err != nil {
		return Data, err
	}

	ArtistData, err := ioutil.ReadAll(resArtist.Body)
	if err != nil {
		return Data, err
	}

	err = json.Unmarshal(ArtistData, &Data)
	if err != nil {
		return Data, err
	}

	return Data, err
}
