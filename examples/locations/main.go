package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/tdabasinskas/go-backstage/backstage"
)

func main() {
	const locationTarget = "https://github.com/tdabasinskas/go-backstage/tree/main/backstage/testdata"

	baseURL, ok := os.LookupEnv("BACKSTAGE_BASE_URL")
	if !ok {
		baseURL = "http://localhost:7007/api/"
	}

	log.Println("Initializing Backstage client...")
	c, _ := backstage.NewClient(baseURL, "default", nil)

	log.Println("Creating new location in a dry-run mode...")
	if location, _, err := c.Catalog.Locations.Create(context.Background(), locationTarget, true); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Location to be created: %v", location)
	}

	log.Println("Creating new location...")
	created, _, err := c.Catalog.Locations.Create(context.Background(), locationTarget, false)
	if err != nil {
		log.Fatal(err)
	} else {
		if created.Location.Target != locationTarget {
			log.Fatalf("Created location target does not match: %v", created.Location.Target)
		} else {
			log.Printf("Location created: %v", created)
		}
	}

	log.Println("Getting created location...")
	if location, _, err := c.Catalog.Locations.GetByID(context.Background(), created.Location.ID); err != nil {
		log.Fatal(err)
	} else {
		if location.Target != locationTarget {
			log.Fatalf("Retrieved location target does not match: %v", location.Target)
		} else {
			log.Printf("Retrieved created location: %v", location)
		}
	}

	log.Println("Listing all locations...")
	if locations, _, err := c.Catalog.Locations.List(context.Background()); err != nil {
		log.Fatal(err)
	} else {
		if len(locations) == 0 {
			log.Fatal("No locations found")
		} else {
			log.Printf("Locations: %v", locations)
		}
	}

	log.Println("Deleting location..")
	if resp, err := c.Catalog.Locations.DeleteByID(context.Background(), created.Location.ID); err != nil {
		log.Fatal(err)
	} else {
		if resp.StatusCode != http.StatusNoContent {
			log.Fatalf("Delete not successful: %s", resp.Status)
		} else {
			log.Printf("Delete response: %v", resp)
		}
	}
}
