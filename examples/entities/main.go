package main

import (
	"context"
	"github.com/tdabasinskas/go-backstage/v2/backstage"
	"log"
	"os"
)

func main() {
	baseURL, ok := os.LookupEnv("BACKSTAGE_BASE_URL")
	if !ok {
		baseURL = "http://localhost:7007/api/"
	}

	log.Println("Initializing Backstage client...")
	c, _ := backstage.NewClient(baseURL, "default", nil)

	log.Println("Getting all entities...")
	if entities, _, err := c.Catalog.Entities.List(context.Background(), nil); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Entities: %v", entities)
	}

	log.Println("Getting component entities...")
	if entities, _, err := c.Catalog.Entities.List(context.Background(), &backstage.ListEntityOptions{
		Filters: []string{"kind=component"},
	}); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Component entities: %v", entities)
	}

	log.Println("Getting location, API and component entities...")
	if entities, _, err := c.Catalog.Entities.List(context.Background(), &backstage.ListEntityOptions{
		Filters: []string{
			"kind=component",
			"kind=location",
			"kind=api",
		},
	}); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Component, location and API entities: %v", entities)
	}

	log.Println("Getting specific component entity by UID...")
	if entity, _, err := c.Catalog.Entities.Get(context.Background(), "06f15cc8-b6b4-4e44-a9bd-1579029f8fb7"); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Component entity: %v", entity)
	}

	log.Println("Getting specific component by name...")
	if component, _, err := c.Catalog.Components.Get(context.Background(), "backstage", ""); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Component: %v", component)
	}
}
