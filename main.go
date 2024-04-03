package main

import (
	"encoding/json"
	"fmt"
	"groupietracker/fyne"
	"io/ioutil"
	"net/http"
)

type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type Geometry struct {
	Location Location `json:"location"`
}

type Result struct {
	Geometry Geometry `json:"geometry"`
}

type GeocodingResponse struct {
	Results []Result `json:"results"`
}

func main() {

	fyne.Begin()

	// URL de l'API Google Maps Geocoding avec une adresse fictive pour l'exemple
	url := "https://google-maps-geocoding.p.rapidapi.com/geocode/json?address=164%20Townsend%20St.%2C%20San%20Francisco%2C%20CA&language=en"

	// Création de la requête HTTP avec l'URL spécifiée
	req, _ := http.NewRequest("GET", url, nil)

	// Ajout des en-têtes requis pour l'API Google Maps Geocoding
	req.Header.Add("X-RapidAPI-Key", "1b86660b93mshe92a2f7cc06d091p1ff205jsnb01707f9c94b")
	req.Header.Add("X-RapidAPI-Host", "google-maps-geocoding.p.rapidapi.com")

	// Envoi de la requête HTTP et récupération de la réponse
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Erreur lors de l'envoi de la requête:", err)
		return
	}
	defer res.Body.Close()

	// Lecture du corps de la réponse
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du corps de la réponse:", err)
		return
	}

	// Déclaration d'une variable pour stocker la réponse de l'API Google Maps Geocoding
	var response GeocodingResponse

	// Décodage du JSON de la réponse dans la variable de réponse déclarée précédemment
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Erreur lors du décodage de la réponse JSON:", err)
		return
	}

	// Vérification si la réponse contient des résultats
	if len(response.Results) > 0 {
		// Affichage des coordonnées géographiques (latitude et longitude) du premier résultat
		lat := response.Results[0].Geometry.Location.Lat
		lng := response.Results[0].Geometry.Location.Lng
		fmt.Println("Coordonnées géographiques (latitude, longitude) :", lat, lng)
	} else {
		fmt.Println("Aucun résultat trouvé.")
	}
}
