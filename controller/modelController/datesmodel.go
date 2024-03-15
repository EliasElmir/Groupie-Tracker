package testmodel

import (
	"encoding/json"
	"fmt"
	"groupietracker/model"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Dates(id int) *model.Dates {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/dates"+"/"+strconv.Itoa(id), nil)
	if err != nil {
		fmt.Print(err.Error())
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}
	var responseObject model.Dates
	json.Unmarshal(bodyBytes, &responseObject)
	return &responseObject
}
