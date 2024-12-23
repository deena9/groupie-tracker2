package main

import (
	"fmt"
	"log"
	"net/http"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Get the artist ID from URL
	// e.g. /artist/1

	FetchAllData()

	artistID := r.URL.Path[len("/artist/"):]
	if artistID == "" {
		http.Error(w, "Artist ID is required", http.StatusBadRequest)
		return
	}

	id := 0
	// Convert the artist ID from string to int
	fmt.Sscanf(artistID, "%d", &id)

	// Fetch the artist and related data
	artist, locations, dates, relations, err := FetchArtistData(id)
	if err != nil {
		http.Error(w, "Failed to fetch artist data", http.StatusInternalServerError)
		return
	}

	stringRelations := []string{}
	for location, dates := range relations {
		for _, date := range dates {
			stringRelations = append(stringRelations, location+" "+date)
		}
	}
	locationCount := 0

	APD := ArtistPageData{
		ID:            artist.ID,
		Image:         artist.Image,
		Name:          artist.Name,
		Members:       artist.Members,
		CreationDate:  artist.CreationDate,
		FirstAlbum:    artist.FirstAlbum,
		Locations:     locations,
		Dates:         dates,
		Relations:     stringRelations,
		LocationCount: locationCount,
	}

	if err := artistTmpl.Execute(w, APD); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Println("Template render error:", err)
	}

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	FetchAllData()

	if err := homeTmpl.Execute(w, artists); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		log.Println("Template render error:", err)
	}
}
