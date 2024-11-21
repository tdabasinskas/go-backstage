package backstage

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

// TestKindLocationGet tests functionality of getting a location.
func TestKindLocationGet(t *testing.T) {
	const dataFile = "testdata/location.json"
	const location = "example"

	expected := LocationEntityV1alpha1{}
	expectedData, _ := os.ReadFile(dataFile)
	err := json.Unmarshal(expectedData, &expected)

	assert.FileExists(t, dataFile, "Test data file should exist")
	assert.NoError(t, err, "Unmarshal should not return an error")

	baseURL, _ := url.Parse("https://foo:1234/api")
	defer gock.Off()
	gock.New(baseURL.String()).
		MatchHeader("Accept", "application/json").
		Get(fmt.Sprintf("/catalog/entities/by-name/location/default/%s", location)).
		Reply(200).
		File(dataFile)

	c, _ := NewClient(baseURL.String(), "", nil)
	s := newLocationService(&entityService{
		client:  c,
		apiPath: "/catalog/entities",
	})

	actual, resp, err := s.Get(context.Background(), location, "")
	assert.NoError(t, err, "Get should not return an error")
	assert.NotEmpty(t, resp, "Response should not be empty")
	assert.EqualValues(t, &expected, actual, "Response body should match the one from the server")
}

// TestKindLocationCreateByID tests functionality of creating a new location.
func TestKindLocationCreateByID(t *testing.T) {
	const dataFile = "testdata/location_create.json"
	const target = "https://github.com/datolabs-io/go-backstage/test"

	expected := LocationCreateResponse{}
	expectedData, _ := os.ReadFile(dataFile)
	err := json.Unmarshal(expectedData, &expected)

	assert.FileExists(t, dataFile, "Test data file should exist")
	assert.NoError(t, err, "Unmarshal should not return an error")

	baseURL, _ := url.Parse("https://foo:1234/api")
	defer gock.Off()
	gock.New(baseURL.String()).
		MatchHeader("Accept", "application/json").
		Post("/catalog/locations").
		MatchParam("dryRun", "false").
		Reply(200).
		JSON(&LocationCreateResponse{
			Location: &LocationResponse{
				ID:     "830d2354-8bbb-42d1-a751-2959f6da5416",
				Type:   "url",
				Target: target,
			},
			Entities: []Entity{},
		})

	c, _ := NewClient(baseURL.String(), "", nil)
	s := newLocationService(&entityService{
		client:  c,
		apiPath: "/catalog/entities",
	})

	actual, resp, err := s.Create(context.Background(), target, false)
	assert.NoError(t, err, "Get should not return an error")
	assert.NotEmpty(t, resp, "Response should not be empty")
	assert.EqualValues(t, &expected, actual, "Response body should match the one from the server")
}

// TestKindLocationCreateByID_DryRun tests functionality of creating a new location.
func TestKindLocationCreateByID_DryRun(t *testing.T) {
	const dataFile = "testdata/location_create_dryrun.json"
	const target = "https://github.com/datolabs-io/go-backstage/test"

	expected := LocationCreateResponse{}
	expectedData, _ := os.ReadFile(dataFile)
	err := json.Unmarshal(expectedData, &expected)

	assert.FileExists(t, dataFile, "Test data file should exist")
	assert.NoError(t, err, "Unmarshal should not return an error")

	baseURL, _ := url.Parse("https://foo:1234/api")
	defer gock.Off()
	gock.New(baseURL.String()).
		MatchHeader("Accept", "application/json").
		Post("/catalog/locations").
		MatchParam("dryRun", "true").
		Reply(200).
		JSON(&LocationCreateResponse{
			Location: &LocationResponse{
				ID:     "830d2354-8bbb-42d1-a751-2959f6da5416",
				Type:   "url",
				Target: target,
			},
			Entities: []Entity{},
		})

	c, _ := NewClient(baseURL.String(), "", nil)
	s := newLocationService(&entityService{
		client:  c,
		apiPath: "/catalog/entities",
	})

	actual, resp, err := s.Create(context.Background(), target, true)
	assert.NoError(t, err, "Get should not return an error")
	assert.NotEmpty(t, resp, "Response should not be empty")
	assert.EqualValues(t, &expected, actual, "Response body should match the one from the server")
}

// TestKindLocationGetByID tests functionality of getting a location by its ID.
func TestKindLocationGetByID(t *testing.T) {
	const dataFile = "testdata/location_by_id.json"
	const id = "830d2354-8bbb-42d1-a751-2959f6da5416"

	expected := LocationResponse{}
	expectedData, _ := os.ReadFile(dataFile)
	err := json.Unmarshal(expectedData, &expected)

	assert.FileExists(t, dataFile, "Test data file should exist")
	assert.NoError(t, err, "Unmarshal should not return an error")

	baseURL, _ := url.Parse("https://foo:1234/api")
	defer gock.Off()
	gock.New(baseURL.String()).
		MatchHeader("Accept", "application/json").
		Get(fmt.Sprintf("/catalog/locations/%s", id)).
		Reply(200).
		File(dataFile)

	c, _ := NewClient(baseURL.String(), "", nil)
	s := newLocationService(&entityService{
		client:  c,
		apiPath: "/catalog/entities",
	})

	actual, resp, err := s.GetByID(context.Background(), id)
	assert.NoError(t, err, "Get should not return an error")
	assert.NotEmpty(t, resp, "Response should not be empty")
	assert.EqualValues(t, &expected, actual, "Response body should match the one from the server")
}

// TestKindLocationList tests functionality of getting all locations.
func TestKindLocationList(t *testing.T) {
	const dataFile = "testdata/locations.json"

	var expected []LocationListResponse
	expectedData, _ := os.ReadFile(dataFile)
	err := json.Unmarshal(expectedData, &expected)

	assert.FileExists(t, dataFile, "Test data file should exist")
	assert.NoError(t, err, "Unmarshal should not return an error")

	baseURL, _ := url.Parse("https://foo:1234/api")
	defer gock.Off()
	gock.New(baseURL.String()).
		MatchHeader("Accept", "application/json").
		Get("/catalog/locations").
		Reply(200).
		File(dataFile)

	c, _ := NewClient(baseURL.String(), "", nil)
	s := newLocationService(&entityService{
		client:  c,
		apiPath: "/catalog/entities",
	})

	actual, resp, err := s.List(context.Background())
	assert.NoError(t, err, "Get should not return an error")
	assert.NotEmpty(t, resp, "Response should not be empty")
	assert.EqualValues(t, expected, actual, "Response body should match the one from the server")
}

// TestEntityServiceDelete tests the deletion of an entity.
func TestKindLocationDeleteByID(t *testing.T) {
	const id = "id"

	baseURL, _ := url.Parse("https://foo:1234/api")
	defer gock.Off()
	gock.New(baseURL.String()).
		MatchHeader("Accept", "application/json").
		Delete(fmt.Sprintf("/catalog/locations/%s", id)).
		Reply(http.StatusNoContent)

	c, _ := NewClient(baseURL.String(), "", nil)
	s := newLocationService(&entityService{
		client:  c,
		apiPath: "/catalog/entities",
	})

	resp, err := s.DeleteByID(context.Background(), id)
	assert.NoError(t, err, "Delete should not return an error")
	assert.NotEmpty(t, resp, "Response should not be empty")
	assert.EqualValues(t, http.StatusNoContent, resp.StatusCode, "Response status code should match the one from the server")
}
