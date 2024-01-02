package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

var theArtists Artists

// var theLocations Locations
// var theDates Dates
var theRelations Relations

func getApi() {
	artistInfo, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("error error!")
	}
	defer artistInfo.Body.Close()
	artistBody, err := io.ReadAll(artistInfo.Body)
	json.Unmarshal(artistBody, &theArtists)
	/* locationInfo, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		fmt.Println("error")
	}
	defer locationInfo.Body.Close()
	locationBody, err := io.ReadAll(locationInfo.Body)
	json.Unmarshal(locationBody, &theLocations)
	dateInfo, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		fmt.Println("error")
	}
	defer dateInfo.Body.Close()
	dateBody, err := io.ReadAll(dateInfo.Body)
	json.Unmarshal(dateBody, &theDates) */
	relationInfo, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		fmt.Println("error")
	}
	defer relationInfo.Body.Close()
	relationBody, err := io.ReadAll(relationInfo.Body)
	json.Unmarshal(relationBody, &theRelations)
	bArray := theRelations.Index
	for i, sa := range theArtists {
		for _, sb := range bArray {
			if sa.Id == sb.Id {
				theArtists[i].DatesLocations = sb.DatesLocations
				break
			}
		}
	}
}
