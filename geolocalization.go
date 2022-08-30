package groupietracker

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"time"
)

const (
	baseUrl = "https://maps.googleapis.com/maps/api/geocode/json?address="
	apiKey  = "&key=AIzaSyA2wftlgG9fePMHDv7JweSNrB3QncemScQ"
)

func getCoord(url string, coordChan chan Coordinate) {
	resp, err := http.Get(url)
	if err != nil {
		// retry?
		log.Fatal("error - get geocode")
	}
	defer resp.Body.Close()
	respData, _ := io.ReadAll(resp.Body)

	var geoLocData GMResponse
	err = json.Unmarshal(respData, &geoLocData)
	// err = json.NewDecoder(resp.Body).Decode(&concertLoc)

	var coord Coordinate
	fmt.Println("geo", geoLocData)
	// if len(geoLocData.Results) == 0 {
	// 	log.Println("API REQUEST DENIED")
	// 	return
	// }
	coord.Lat = geoLocData.Results[0].Geometry.Location.Lat
	coord.Lng = geoLocData.Results[0].Geometry.Location.Lng
	coordChan <- coord
	// return coord
}

func buildCoordDate(coord Coordinate, dateStr string) CoordDate {
	var coDate CoordDate

	coDate.Coord = coord

	dateInTime, _ := time.Parse("02-01-2006", dateStr)
	coDate.Date = dateInTime

	return coDate
}

func RequestArtistGeocodes(artist *ArtistAllData) {
	for address, dates := range artist.DatesLocations {
		url := baseUrl + address + apiKey

		coordChan := make(chan Coordinate)
		go getCoord(url, coordChan)
		co := <-coordChan
		for d := 0; d < len(dates); d++ {
			artist.CoordsDates = append(artist.CoordsDates, buildCoordDate(co, dates[d]))
		}
	}
	sort.Slice(artist.CoordsDates, func(i, j int) bool { return artist.CoordsDates[j].Date.After(artist.CoordsDates[i].Date) })
}
