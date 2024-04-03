package DataAPI

import (
	"strings"
)

// CreationDateFilter filter in ascending order the creation date of all groups
func CreationDateFilter() (Artist, error) {
	Data, err := GetArtistData(false)
	if err != nil {
		return Data, err
	}
	for i := 0; i < len(Data); i++ {
		for j := 0; j < len(Data)-1; j++ {
			if Data[j].CreationDate > Data[j+1].CreationDate {
				prevData := Data[j+1]
				Data[j+1] = Data[j]
				Data[j] = prevData
			}
		}
	}

	return Data, err
}

// FirstAlbumFilter filter in ascending number the date of the first album
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

// NbOfMemberFilter filter in ascending order the number of member of a group
func NbOfMemberFilter() (Artist, error) {
	Data, err := GetArtistData(false)
	if err != nil {
		return Data, err
	}

	for i := 0; i < len(Data); i++ {
		for j := 0; j < len(Data)-1; j++ {
			if len(Data[j].Members) > len(Data[j+1].Members) {
				prevData := Data[j+1]
				Data[j+1] = Data[j]
				Data[j] = prevData
			}
		}
	}

	return Data, err
}

// LocationFilter filter the artist depending on the location in the parameter
/*func LocationFilter() (Artist, error) {
	Data, err := GetArtistData(true)
	if err != nil {
		return Data, err
	}

	for i := 0; i < len(Data); i++ {
		for ConcertLocation, _ := range Data[i].Relations {

		}
	}

	return Data, err
}*/

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
