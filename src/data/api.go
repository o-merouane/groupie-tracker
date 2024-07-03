package data

import (
	"encoding/json"
	"net/http"
)

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type LocationIndex struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type DatesIndex struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func FetchArtistData() ([]Artist, error) {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	var artists []Artist
	err = json.NewDecoder(rsp.Body).Decode(&artists)
	if err != nil {
		return nil, err
	}
	return artists, nil
}

func FetchLocationIndexData() ([]LocationIndex, error) {
	url := "https://groupietrackers.herokuapp.com/api/locations"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var locations []LocationIndex
	err = json.NewDecoder(resp.Body).Decode(&locations)
	if err != nil {
		return nil, err
	}

	return locations, nil
}

func FetchDatesIndexData() ([]DatesIndex, error) {
	url := "https://groupietrackers.herokuapp.com/api/dates"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var dates []DatesIndex
	err = json.NewDecoder(resp.Body).Decode(&dates)
	if err != nil {
		return nil, err
	}

	return dates, nil
}

func FetchRelationsData() (Relations, error) {
	url := "https://groupietrackers.herokuapp.com/api/relation"
	resp, err := http.Get(url)
	if err != nil {
		return Relations{}, err
	}
	defer resp.Body.Close()

	var relations Relations
	err = json.NewDecoder(resp.Body).Decode(&relations)
	if err != nil {
		return Relations{}, err
	}

	return relations, nil
}
