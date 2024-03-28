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
	DateData, err := GetDateData()
	LocationData, err := GetLocationData()
	DataRelation, err := GetRelationsData()
	resArtist, err := http.Get(URL + "/artists")
	if err != nil {
		return Data, err
	}

	ArtistData, err := ioutil.ReadAll(resArtist.Body)
	if err != nil {
		return Data, err
	}

	err = json.Unmarshal(ArtistData, &Data)

	for i, _ := range Data {
		Data[i].ConcertDates = DateData.Dates[i].Dates
		Data[i].Locations = LocationData.Location[i].Locations
		Data[i].Relations = make(map[string][]string)
		for key, value := range DataRelation.Relation[i].DatesLocations {
			Data[i].Relations[key] = value
		}
	}

	return Data, err
}

func GetLocationData() (LocationIndex, error) {
	var Data LocationIndex
	resData, err := http.Get(URL + "/locations")
	if err != nil {
		return Data, err
	}

	DateData, err := ioutil.ReadAll(resData.Body)
	if err != nil {
		return Data, err
	}

	err = json.Unmarshal(DateData, &Data)
	if err != nil {
		return Data, nil
	}
	return Data, err
}

func GetDateData() (DatesIndex, error) {
	var Data DatesIndex
	resData, err := http.Get(URL + "/dates")
	if err != nil {
		return Data, err
	}

	DateData, err := ioutil.ReadAll(resData.Body)
	if err != nil {
		return Data, err
	}

	err = json.Unmarshal(DateData, &Data)
	if err != nil {
		return Data, nil
	}
	return Data, err
}

func GetRelationsData() (RelationIndex, error) {
	var Data RelationIndex
	resData, err := http.Get(URL + "/relation")
	if err != nil {
		return Data, err
	}

	RelationData, err := ioutil.ReadAll(resData.Body)
	if err != nil {
		return Data, err
	}

	err = json.Unmarshal(RelationData, &Data)
	if err != nil {
		return Data, err
	}

	return Data, err
}

func GetDateByID(id int) DataDate {
	var Data DataDate
	resData, _ := http.Get(URL + "/dates/" + string(id))

	RelationData, _ := ioutil.ReadAll(resData.Body)

	_ = json.Unmarshal(RelationData, &Data)

	return Data
}

func GetArtistByID(id int) DataArtist {
	var Data DataArtist
	resData, _ := http.Get(URL + "/artists/" + string(id))

	RelationData, _ := ioutil.ReadAll(resData.Body)

	_ = json.Unmarshal(RelationData, &Data)

	return Data
}
