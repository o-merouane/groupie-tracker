package data

import (
	"encoding/json"
	"fmt"
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

type RelationApiResponse struct {
	Index []Relations `json:"Index"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type DatesApiResponse struct {
	Index []Dates `json:"index"`
}

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type LocationAPIResponse struct {
	Index []LocationItem `json:"index"`
}

type LocationItem struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
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

func FetchLocationData(url string) (*LocationAPIResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse LocationAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &apiResponse, nil
}

func FetchDatesData(url string) (*DatesApiResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponse DatesApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &apiResponse, nil
}

func FetchRelationsData(url string) (*RelationApiResponse, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer rsp.Body.Close()

	var apiResponse RelationApiResponse
	if err := json.NewDecoder(rsp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}
	return &apiResponse, nil
}

type CombinedArtistData struct {
	Artist    Artist
	Locations []string
	Dates     []string
}

func FetchCombinedArtistData() ([]CombinedArtistData, error) {
	artists, err := FetchArtistData()
	if err != nil {
		return nil, err
	}

	locations, err := FetchLocationData("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}

	dates, err := FetchDatesData("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}

	var combinedData []CombinedArtistData

	for _, artist := range artists {
		var artistLocations []string
		var artistDates []string

		for _, locationItem := range locations.Index {
			if locationItem.ID == artist.ID {
				artistLocations = locationItem.Locations
			}
		}

		for _, dateItem := range dates.Index {
			if dateItem.ID == artist.ID {
				artistDates = dateItem.Dates
			}
		}

		combinedData = append(combinedData, CombinedArtistData{
			Artist:    artist,
			Locations: artistLocations,
			Dates:     artistDates,
		})
	}

	return combinedData, nil
}
