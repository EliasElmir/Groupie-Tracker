/*package DataAPI

import (
	"strconv"
	"strings"
)

// CreationDateFilter filter and return artist corresponding to the parameter of the creation date of the group/artist
func CreationDateFilter(year string) (Artist, error) {
	Data, err := GetArtistData(false)
	if err != nil {
		return Data, err
	}
	var FilteredData Artist
	for i := 0; i < len(Data); i++ {
		if strings.Contains(strconv.Itoa(Data[i].CreationDate), year) {
			FilteredData = append(FilteredData, Data[i])
		}
	}

	return FilteredData, err
}

// FirstAlbumFilter filter the artist from the year of the first album in the parameter
func FirstAlbumFilter(year string) (Artist, error) {
	Data, err := GetArtistData(false)
	if err != nil {
		return Data, err
	}
	var FilteredData Artist
	for i := 0; i < len(Data); i++ {
		if strings.Contains(Data[i].FirstAlbum, year) {
			FilteredData = append(FilteredData, Data[i])
		}
	}

	return FilteredData, err
}

// NbOfMemberFilter filter the artist depending on the number of member of a group/artist
func NbOfMemberFilter(nbfilter int) (Artist, error) {
	Data, err := GetArtistData(false)
	if err != nil {
		return Data, err
	}
	var FilteredData Artist
	for i := 0; i < len(Data); i++ {
		if nbfilter == len(Data[i].Members) {
			FilteredData = append(FilteredData, Data[i])
		}
	}

	return FilteredData, err
}

// LocationFilter filter the artist depending on the location in the parameter
func LocationFilter(location string) (Artist, error) {
	Data, err := GetArtistData(true)
	if err != nil {
		return Data, err
	}

	var FilteredData Artist

	for i := 0; i < len(Data); i++ {
		for ConcertLocation, _ := range Data[i].Relations {
			if strings.Contains(strings.ToLower(ConcertLocation), strings.ToLower(location)) {
				FilteredData = append(FilteredData, Data[i])
			}
		}
	}

	return FilteredData, err
}
