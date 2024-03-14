package model

type Relation struct {
	Id             int64               `json "id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
