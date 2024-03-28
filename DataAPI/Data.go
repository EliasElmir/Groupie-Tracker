package DataAPI

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

const URL = "https://groupietrackers.herokuapp.com/api"

type Artist []DataArtist
type LocationIndex struct {
	Location []DataLocation `json:"index"`
}

type DatesIndex struct {
	Dates []DataDate `json:"index"`
}

type RelationIndex struct {
	Relation []DataRelations `json:"index"`
}

type DataArtist struct {
	Id              int      `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    int      `json:"creationDate"`
	FirstAlbum      string   `json:"firstAlbum"`
	LocationURL     string   `json:"locations"`
	ConcertDatesURL string   `json:"concertDates"`
	RelationURL     string   `json:"relations"`
	Locations       []string
	ConcertDates    []string
	Relations       map[string][]string
}

type DataLocation struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"`
	Dates     []string
}

type DataDate struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DataRelations struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetArtistData(isAllDataNeeded bool) (Artist, error) {
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
	if isAllDataNeeded {
		DateData, err := GetDateData()
		LocationData, err := GetLocationData()
		DataRelation, err := GetRelationsData()
		if err != nil {
			return nil, err
		}
		for i, _ := range Data {
			Data[i].ConcertDates = DateData.Dates[i].Dates
			Data[i].Locations = LocationData.Location[i].Locations
			Data[i].Relations = make(map[string][]string)
			for key, value := range DataRelation.Relation[i].DatesLocations {
				Data[i].Relations[key] = value
			}
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
		return Data, err
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

func GetDateByID(id int) (DataDate, error) {
	var Data DataDate
	resData, err := http.Get(URL + "/dates/" + strconv.Itoa(id))
	if err != nil {
		return Data, err
	}

	DateData, err := ioutil.ReadAll(resData.Body)
	if err != nil {
		return Data, err
	}

	err = json.Unmarshal(DateData, &Data)
	if err != nil {
		return Data, err
	}

	return Data, err
}

func GetArtistByID(id int) (DataArtist, error) {
	var Data DataArtist
	resData, err := http.Get(URL + "/artists/" + strconv.Itoa(id))
	if err != nil {
		return Data, err
	}

	ArtistData, err := ioutil.ReadAll(resData.Body)
	if err != nil {
		return Data, err
	}

	err = json.Unmarshal(ArtistData, &Data)
	if err != nil {
		return Data, err
	}

	return Data, err
}

func GetLocationByID(id int) (DataLocation, error) {
	var Data DataLocation
	resData, err := http.Get(URL + "/locations/" + strconv.Itoa(id))
	if err != nil {
		return Data, err
	}

	LocationData, err := ioutil.ReadAll(resData.Body)
	if err != nil {
		return Data, err
	}

	err = json.Unmarshal(LocationData, &Data)
	if err != nil {
		return Data, err
	}

	DatesData, err := GetDateByID(id)
	Data.Dates = DatesData.Dates

	return Data, err
}

func GetRelationByID(id int) (DataRelations, error) {
	var Data DataRelations
	resData, err := http.Get(URL + "/relation/" + strconv.Itoa(id))
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
