package DataAPI

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
func FirstAlbumFilter() (Artist, error) {
	Data, err := GetArtistData(false)
	if err != nil {
		return Data, err
	}

	Dates := make([][]string, len(Data))
	for i := 0; i < len(Data); i++ {
		Dates[i] = strings.SplitN(Data[i].FirstAlbum, "-", -1)
	}

	NewArtistOrder := sortDates(Data, Dates)

	return NewArtistOrder, err
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

// sortDates Sort dates in ascending order
func sortDates(Data Artist, Dates [][]string) Artist {
	// Loop to sort by year
	for i := 0; i < len(Data)-1; i++ {
		for j := 0; j < len(Data)-1; j++ {
			if Dates[j][2] > Dates[j+1][2] {
				prevDate := Dates[j+1]
				Dates[j+1] = Dates[j]
				Dates[j] = prevDate

				prevData := Data[j+1]
				Data[j+1] = Data[j]
				Data[j] = prevData
			}
		}
	}

	// Loop to sort by month
	for i := 0; i < len(Data)-1; i++ {
		for j := 0; j < len(Data)-1; j++ {
			if Dates[j][2] == Dates[j+1][2] {
				if Dates[j][1] > Dates[j+1][1] {
					prevDate := Dates[j+1]
					Dates[j+1] = Dates[j]
					Dates[j] = prevDate

					prevData := Data[j+1]
					Data[j+1] = Data[j]
					Data[j] = prevData
				}
			}
		}
	}

	// Loop to sort by day
	for i := 0; i < len(Data)-1; i++ {
		for j := 0; j < len(Data)-1; j++ {
			if Dates[j][2] == Dates[j+1][2] {
				if Dates[j][1] > Dates[j+1][1] {
					if Dates[j][0] > Dates[j+1][0] {
						prevDate := Dates[j+1]
						Dates[j+1] = Dates[j]
						Dates[j] = prevDate

						prevData := Data[j+1]
						Data[j+1] = Data[j]
						Data[j] = prevData
					}
				}
			}
		}
	}

	return Data
}
