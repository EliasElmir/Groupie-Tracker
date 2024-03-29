package API

import (
	"encoding/json"
	"fmt"
	model "groupietracker/Structure"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Location(id int) *model.Locations {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://groupietrackers.herokuapp.com/api/locations"+"/"+strconv.Itoa(id), nil)
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
	var responseObject model.Locations
	json.Unmarshal(bodyBytes, &responseObject)
	fmt.Println(responseObject.Location)
	return &responseObject
}
