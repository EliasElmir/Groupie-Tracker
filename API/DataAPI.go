package DataAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"
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

type RelationSTRUCT struct {
	Id             int64               `json "id"`
	DatesLocations map[string][]string `json:"datesLocations"`
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

type Locations struct {
	Id       int64    `json:"id"`
	Location []string `json:"locations"`
}

func GetArtistData(isAllDataNeeded bool) (Artist, error) {
	// GetArtistData get every artist information from the api, with a bool to know if it needs to get locations "concertdates" and relations as well
	var Data Artist

	resArtist, err := http.Get(URL + "/artists")
	if err != nil {
		return Data, err
	}

	ArtistData, err := io.ReadAll(resArtist.Body)
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
	// GetLocationData get the location of concert of every artist of the api
	var Data LocationIndex
	resData, err := http.Get(URL + "/locations")
	if err != nil {
		return Data, err
	}

	DateData, err := io.ReadAll(resData.Body)
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
	// GetDateData get the date of every artist of the api
	var Data DatesIndex
	resData, err := http.Get(URL + "/dates")
	if err != nil {
		return Data, err
	}

	DateData, err := io.ReadAll(resData.Body)
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
	// GetRelationsData get the relation of every artist of the api
	var Data RelationIndex
	resData, err := http.Get(URL + "/relation")
	if err != nil {
		return Data, err
	}

	RelationData, err := io.ReadAll(resData.Body)
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
	// GetDateByID get the date of a specific artist
	var Data DataDate
	resData, err := http.Get(URL + "/dates/" + strconv.Itoa(id))
	if err != nil {
		return Data
	}

	DateData, err := io.ReadAll(resData.Body)
	if err != nil {
		return Data
	}

	err = json.Unmarshal(DateData, &Data)
	if err != nil {
		return Data
	}

	return Data
}

func GetArtistByID(id int) DataArtist {
	// GetArtistByID get the information of a specific artist with the id
	var Data DataArtist
	resData, err := http.Get(URL + "/artists/" + strconv.Itoa(id))
	if err != nil {
		return Data
	}

	ArtistData, err := io.ReadAll(resData.Body)
	if err != nil {
		return Data
	}

	err = json.Unmarshal(ArtistData, &Data)
	if err != nil {
		return Data
	}

	return Data
}

func GetLocationByID(id int) DataLocation {
	// GetLocationByID get the location of a specific artist with the id
	var Data DataLocation
	resData, err := http.Get(URL + "/locations/" + strconv.Itoa(id))
	if err != nil {
		return Data
	}

	LocationData, err := io.ReadAll(resData.Body)
	if err != nil {
		return Data
	}

	err = json.Unmarshal(LocationData, &Data)
	if err != nil {
		return Data
	}

	DatesData := GetDateByID(id)
	Data.Dates = DatesData.Dates

	return Data
}

func GetRelationByID(id int) DataRelations {
	// GetRelationByID get the relation of a specific artist with the id
	var Data DataRelations
	resData, err := http.Get(URL + "/relation/" + strconv.Itoa(id))
	if err != nil {
		return Data
	}

	RelationData, err := io.ReadAll(resData.Body)
	if err != nil {
		return Data
	}

	err = json.Unmarshal(RelationData, &Data)
	if err != nil {
		return Data
	}

	return Data
}

func Location(id int) *Locations {
	// Location get the location of a specific artist with the id
	url := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%d", id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP:", err)
		return nil
	}
	defer resp.Body.Close()

	var responseObject Locations
	if err := json.NewDecoder(resp.Body).Decode(&responseObject); err != nil {
		fmt.Println("Erreur lors du décodage de la réponse JSON:", err)
		return nil
	}

	fmt.Println(responseObject.Location)
	return &responseObject
}

func Relation(id int) RelationSTRUCT {
	// Relation get the relation of a specific artist with the id
	var relations RelationSTRUCT

	err := get("https://groupietrackers.herokuapp.com/api/relation/"+strconv.Itoa(id), &relations)
	if err != nil {
		return relations
	}
	return relations
}

func get(url string, target interface{}) error {
	// get get the data from the url and put it in the target
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(target)

	if err != nil {
		return err
	}

	return nil
}

func Search(query string) ([]DataArtist, error) {
	// Search get the artist that match the query
	allArtists, err := GetArtistData(true)
	if err != nil {
		return nil, err
	}

	var searchResults []DataArtist

	for _, artist := range allArtists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(query)) {
			searchResults = append(searchResults, artist)
		}
	}

	return searchResults, nil
}

func SortByCreationDateAscending(artists Artist) Artist {
	sort.Slice(artists, func(i, j int) bool {
		return artists[i].CreationDate < artists[j].CreationDate
	})
	return artists
}

func SortByCreationDateAndFirstAlbum(artists Artist) Artist {
	sort.Slice(artists, func(i, j int) bool {
		if artists[i].CreationDate != artists[j].CreationDate {
			return artists[i].CreationDate < artists[j].CreationDate
		}
		// Si les dates de création sont égales, trie par date du premier album
		return artists[i].FirstAlbum < artists[j].FirstAlbum
	})
	return artists
}

func SortByNumberOfMembersDescending(artists Artist) Artist {
	sort.Slice(artists, func(i, j int) bool {
		return len(artists[i].Members) > len(artists[j].Members)
	})
	return artists
}
