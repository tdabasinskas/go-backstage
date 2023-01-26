package main

import (
	"context"
	"log"
	"os"

	"github.com/tdabasinskas/go-backstage/backstage"
)

func main() {
	const locationTarget = "https://github.com/tdabasinskas/backstage-go/test"

	baseURL, ok := os.LookupEnv("BACKSTAGE_BASE_URL")
	if !ok {
		baseURL = "http://localhost:7007/api/"
	}

	log.Println("Initializing Backstage client...")
	c, _ := backstage.NewClient(baseURL, "default", nil)

	log.Println("Creating new location in a dry-run mode...")
	if location, _, err := c.Catalog.Locations.Create(context.Background(), locationTarget, false); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Location to be created: %v", location)
	}

	log.Println("Creating new location...")
	created, _, err := c.Catalog.Locations.Create(context.Background(), locationTarget, true)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Location created: %v", created)
	}

	log.Println("Getting created location...")
	if location, _, err := c.Catalog.Locations.GetByID(context.Background(), created.Location.ID); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Retrieved created location: %v", location)
	}

	log.Println("Listing all locations...")
	if locations, _, err := c.Catalog.Locations.List(context.Background()); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Locations: %v", locations)
	}

	log.Println("Deleting location..")
	if resp, err := c.Catalog.Locations.DeleteByID(context.Background(), created.Location.ID); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Delete response: %v", resp)
	}
}
