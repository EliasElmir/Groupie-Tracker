package testmodel

import (
	"groupietracker/model"
	"strconv"
)

func Relation(id int) model.Relation {
	var relations model.Relation

	err := get("https://groupietrackers.herokuapp.com/api/relation/"+strconv.Itoa(id), &relations)
	if err != nil {
		return relations
	}
	return relations
}
