package groupietracker

import "time"

type ArtistAllData struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Locations      []string            `json:"locations"`
	ConcertDates   []string            `json:"concertDates"`
	DatesLocations map[string][]string `json:"datesLocations"`
	CoordsDates    []CoordDate
}

type Artistforhomepage struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ArtistDataWithRelation struct {
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

type LocationAllData struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

type RelationAllData struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type DateAllData struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

type DateRemoveIndex struct {
	Index []DateAllData `json:"index"`
}

type LocationRemoveIndex struct {
	Index []LocationAllData `json:"index"`
}

type RelationRemoveIndex struct {
	Index []RelationAllData `json:"index"`
}

var (
	ArtistsFull []ArtistAllData
	Artists     []ArtistDataWithRelation
	Dates       DateRemoveIndex
	Locations   LocationRemoveIndex
	Relations   RelationRemoveIndex
	Test        ArtistAllData
)

// geolocalization
type CoordDate struct {
	Coord Coordinate
	Date  time.Time
}
type Coordinate struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
type GMResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64
				Lng float64
			}
		}
	}
	Status string
}
